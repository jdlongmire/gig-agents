package server

import (
	"context"
	"database/sql"
	"fmt"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/bobbyhiddn/CrewPort/internal/auth"
	"github.com/bobbyhiddn/CrewPort/internal/config"
	dbsqlc "github.com/bobbyhiddn/CrewPort/internal/db/sqlc"
)

// Server wraps the HTTP server and its dependencies.
type Server struct {
	cfg     *config.Config
	db      *sql.DB
	queries *dbsqlc.Querier
	httpSrv *http.Server
}

// New constructs and configures the HTTP server, registering all routes.
func New(cfg *config.Config, db *sql.DB) *Server {
	s := &Server{
		cfg:     cfg,
		db:      db,
		queries: dbsqlc.New(db),
	}
	s.httpSrv = &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.Port),
		Handler:      s.buildRouter(),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	return s
}

// buildRouter creates the chi router with all middleware and routes attached.
func (s *Server) buildRouter() http.Handler {
	r := chi.NewRouter()

	// --- Global middleware ---
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(slogMiddleware)
	r.Use(middleware.Recoverer)

	// OAuth handler (no auth required)
	oauthHandler := auth.NewOAuthHandler(
		s.cfg.GitHubClientID,
		s.cfg.GitHubClientSecret,
		s.cfg.BaseURL,
		s.cfg.JWTSecret,
		s.cfg.AllowedOperators,
		s.db,
	)

	// --- Auth routes (no JWT required) ---
	r.Get("/auth/github", oauthHandler.HandleRedirect)
	r.Get("/auth/github/callback", oauthHandler.HandleCallback)

	// --- API routes ---
	r.Route("/api", func(r chi.Router) {
		r.Get("/health", healthHandler)

		// Protected API routes
		r.Group(func(r chi.Router) {
			r.Use(auth.Middleware(s.cfg.JWTSecret))
			r.Get("/auth/me", meHandler(s.queries))
			r.Post("/auth/logout", logoutHandler)
		})
	})

	// --- SPA fallback: serve web/dist/ with index.html for unknown routes ---
	r.Get("/*", spaHandler())

	return r
}

// spaHandler serves the built Svelte SPA from web/dist/.
// Falls back to index.html for any path not found (client-side routing).
func spaHandler() http.HandlerFunc {
	distDir := "web/dist"
	fsys := os.DirFS(distDir)

	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if path == "/" || path == "" {
			path = "index.html"
		} else {
			// Strip leading slash
			if len(path) > 1 && path[0] == '/' {
				path = path[1:]
			}
		}

		// Check if the file exists in the dist directory
		if _, err := fs.Stat(fsys, path); err == nil {
			http.ServeFileFS(w, r, fsys, path)
			return
		}

		// Fall back to index.html for SPA client-side routing
		http.ServeFileFS(w, r, fsys, "index.html")
	}
}

// Run starts the HTTP server and blocks until a shutdown signal is received.
func (s *Server) Run() error {
	// Channel to listen for OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Start server in background
	serverErr := make(chan error, 1)
	go func() {
		slog.Info("server starting", "addr", s.httpSrv.Addr)
		if err := s.httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			serverErr <- err
		}
	}()

	// Wait for signal or error
	select {
	case err := <-serverErr:
		return fmt.Errorf("server error: %w", err)
	case sig := <-quit:
		slog.Info("shutdown signal received", "signal", sig)
	}

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.httpSrv.Shutdown(ctx); err != nil {
		return fmt.Errorf("graceful shutdown failed: %w", err)
	}

	slog.Info("server stopped cleanly")
	return nil
}

// slogMiddleware logs each request using slog.
func slogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		start := time.Now()
		next.ServeHTTP(ww, r)
		slog.Info("request",
			"method", r.Method,
			"path", r.URL.Path,
			"status", ww.Status(),
			"bytes", ww.BytesWritten(),
			"duration", time.Since(start),
			"request_id", middleware.GetReqID(r.Context()),
		)
	})
}
