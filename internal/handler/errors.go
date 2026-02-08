package handler

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/grovecj/warzone-stats-tracker/internal/codclient"
)

type apiError struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// writeAPIError maps codclient sentinel errors to HTTP status codes and writes a JSON response.
func writeAPIError(w http.ResponseWriter, err error) {
	var (
		status int
		code   string
		msg    string
	)

	switch {
	case errors.Is(err, codclient.ErrPlayerNotFound):
		status = http.StatusNotFound
		code = "player_not_found"
		msg = "Player not found"
	case errors.Is(err, codclient.ErrPrivateProfile):
		status = http.StatusForbidden
		code = "private_profile"
		msg = "Player profile is set to private"
	case errors.Is(err, codclient.ErrTokenExpired):
		status = http.StatusServiceUnavailable
		code = "service_unavailable"
		msg = "CoD API authentication expired"
	case errors.Is(err, codclient.ErrAPIUnavailable):
		status = http.StatusServiceUnavailable
		code = "service_unavailable"
		msg = "CoD API is currently unavailable"
	case errors.Is(err, codclient.ErrRateLimited):
		status = http.StatusTooManyRequests
		code = "rate_limited"
		msg = "Too many requests to CoD API"
	default:
		slog.Error("unhandled error in API handler", "error", err)
		status = http.StatusInternalServerError
		code = "internal_error"
		msg = "An unexpected error occurred"
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(apiError{Error: code, Message: msg})
}
