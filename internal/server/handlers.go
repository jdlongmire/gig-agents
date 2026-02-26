package server

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/bobbyhiddn/CrewPort/internal/auth"
	dbsqlc "github.com/bobbyhiddn/CrewPort/internal/db/sqlc"
)

// healthHandler returns a simple JSON health check response.
func healthHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{
		"status":  "ok",
		"service": "crewport",
	})
}

// meHandler returns the currently-authenticated user's profile.
func meHandler(queries *dbsqlc.Querier) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, ok := auth.ClaimsFromContext(r.Context())
		if !ok {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		account, err := queries.GetAccountByID(r.Context(), claims.Subject)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				http.Error(w, "account not found", http.StatusNotFound)
				return
			}
			slog.Error("me: get account failed", "err", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		writeJSON(w, http.StatusOK, account)
	}
}

// logoutHandler clears the auth cookie and returns 200.
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	auth.ClearTokenCookie(w)
	writeJSON(w, http.StatusOK, map[string]string{"message": "logged out"})
}

// writeJSON encodes v as JSON and writes it with the given status code.
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		slog.Error("writeJSON encode failed", "err", err)
	}
}
