package main

import (
	"context"
	"fmt"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/grovecj/warzone-stats-tracker/internal/cache"
	"github.com/grovecj/warzone-stats-tracker/internal/codclient"
	"github.com/grovecj/warzone-stats-tracker/internal/config"
	"github.com/grovecj/warzone-stats-tracker/internal/database"
	"github.com/grovecj/warzone-stats-tracker/internal/handler"
	"github.com/grovecj/warzone-stats-tracker/internal/router"
	"github.com/grovecj/warzone-stats-tracker/web"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		slog.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: cfg.LogLevel(),
	}))
	slog.SetDefault(logger)

	// Database
	ctx := context.Background()
	pool, err := database.Connect(ctx, cfg.DatabaseURL)
	if err != nil {
		slog.Error("failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer pool.Close()
	slog.Info("database connected")

	// Run migrations
	if err := database.RunMigrations(cfg.DatabaseURL, "migrations"); err != nil {
		slog.Error("failed to run migrations", "error", err)
		os.Exit(1)
	}

	// CoD API client with caching
	codAPI := codclient.New(cfg.CodAPIBaseURL, cfg.CodSSOToken)
	cachedAPI := cache.New(codAPI, cache.DefaultConfig())

	// Static files — use embedded FS in production, nil in dev (Vite proxy handles it)
	var staticFS fs.FS
	if dist, err := fs.Sub(web.DistFS, "dist"); err == nil {
		if _, err := fs.Stat(dist, "index.html"); err == nil {
			staticFS = dist
			slog.Info("serving embedded frontend")
		}
	}

	// Handlers
	adminHandler := handler.NewAdminHandler(cachedAPI)

	// Router
	rawOrigins := strings.Split(cfg.CORSAllowedOrigins, ",")
	var origins []string
	for _, o := range rawOrigins {
		if trimmed := strings.TrimSpace(o); trimmed != "" {
			origins = append(origins, trimmed)
		}
	}
	mux := router.New(origins, staticFS, router.Deps{
		AdminHandler: adminHandler,
		AdminAPIKey:  cfg.AdminAPIKey,
	})

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Graceful shutdown
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)

	go func() {
		slog.Info("server starting", "port", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("server failed", "error", err)
			os.Exit(1)
		}
	}()

	<-done
	slog.Info("shutting down server")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		slog.Error("server shutdown failed", "error", err)
		os.Exit(1)
	}

	slog.Info("server stopped")
}
