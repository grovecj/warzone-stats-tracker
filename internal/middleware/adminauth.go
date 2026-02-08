package middleware

import (
	"encoding/json"
	"net/http"
	"strings"
)

// AdminAuth returns middleware that validates the ADMIN_API_KEY via the Authorization header.
func AdminAuth(apiKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if apiKey == "" {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusServiceUnavailable)
				json.NewEncoder(w).Encode(map[string]string{
					"error":   "admin_not_configured",
					"message": "Admin API key is not configured",
				})
				return
			}

			auth := r.Header.Get("Authorization")
			token := strings.TrimPrefix(auth, "Bearer ")
			if token == "" || token == auth || token != apiKey {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{
					"error":   "unauthorized",
					"message": "Invalid or missing admin API key",
				})
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
