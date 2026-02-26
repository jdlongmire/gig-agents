package main

import (
	"log/slog"
	"os"

	"github.com/bobbyhiddn/CrewPort/internal/config"
	"github.com/bobbyhiddn/CrewPort/internal/db"
	"github.com/bobbyhiddn/CrewPort/internal/server"
)

func main() {
	// Structured JSON logging to stdout
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})))

	cfg, err := config.Load()
	if err != nil {
		slog.Error("config error", "err", err)
		os.Exit(1)
	}

	database, err := db.Open(cfg.DBPath)
	if err != nil {
		slog.Error("database error", "err", err)
		os.Exit(1)
	}
	defer database.Close()

	srv := server.New(cfg, database)
	if err := srv.Run(); err != nil {
		slog.Error("server error", "err", err)
		os.Exit(1)
	}
}
