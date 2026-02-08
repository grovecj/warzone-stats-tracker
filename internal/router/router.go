package router

import (
	"io/fs"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"

	"github.com/grovecj/warzone-stats-tracker/internal/handler"
	"github.com/grovecj/warzone-stats-tracker/internal/middleware"
)

// Deps holds dependencies injected into the router.
type Deps struct {
	AdminHandler  *handler.AdminHandler
	PlayerHandler *handler.PlayerHandler
	MatchHandler  *handler.MatchHandler
	AdminAPIKey   string
}

func New(allowedOrigins []string, staticFS fs.FS, deps Deps) http.Handler {
	r := chi.NewRouter()

	// Global middleware
	r.Use(middleware.Recovery)
	r.Use(middleware.RequestLogger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-Requested-With"},
		ExposedHeaders:   []string{"X-Request-ID", "X-Cache", "X-Data-Age"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	r.Use(middleware.NewRateLimiter(100, 1*time.Minute).Handler)

	// API routes
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", handler.Health)

		// Player routes
		r.Route("/players", func(r chi.Router) {
			if deps.PlayerHandler != nil {
				r.Get("/search", deps.PlayerHandler.SearchPlayer)
				r.Get("/{platform}/{gamertag}/stats", deps.PlayerHandler.GetStats)
			} else {
				r.Get("/search", handler.NotImplemented)
				r.Get("/{platform}/{gamertag}/stats", handler.NotImplemented)
			}
			if deps.MatchHandler != nil {
				r.Get("/{platform}/{gamertag}/matches", deps.MatchHandler.GetMatches)
			} else {
				r.Get("/{platform}/{gamertag}/matches", handler.NotImplemented)
			}
		})

		// Comparison routes (issue #14)
		r.Get("/compare", handler.NotImplemented)

		// Admin routes (protected by ADMIN_API_KEY)
		r.Route("/admin", func(r chi.Router) {
			r.Use(middleware.AdminAuth(deps.AdminAPIKey))
			if deps.AdminHandler != nil {
				r.Post("/token", deps.AdminHandler.UpdateToken)
			}
		})

		// Squad routes (issue #16)
		r.Route("/squads", func(r chi.Router) {
			r.Post("/", handler.NotImplemented)
			r.Get("/{squadID}", handler.NotImplemented)
			r.Put("/{squadID}", handler.NotImplemented)
			r.Delete("/{squadID}", handler.NotImplemented)
			r.Post("/{squadID}/members", handler.NotImplemented)
			r.Delete("/{squadID}/members/{playerID}", handler.NotImplemented)
			r.Get("/{squadID}/stats", handler.NotImplemented)
		})
	})

	// Serve static frontend files with SPA fallback
	if staticFS != nil {
		fileServer := http.FileServerFS(staticFS)
		r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
			// Try to serve the file directly
			path := strings.TrimPrefix(r.URL.Path, "/")
			if path == "" {
				path = "index.html"
			}

			if _, err := fs.Stat(staticFS, path); err != nil {
				// File doesn't exist — serve index.html for SPA routing
				r.URL.Path = "/"
			}

			fileServer.ServeHTTP(w, r)
		})
	}

	return r
}
