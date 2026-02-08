package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/grovecj/warzone-stats-tracker/internal/service"
)

// PlayerHandler holds dependencies for player endpoints.
type PlayerHandler struct {
	playerService *service.PlayerService
}

// NewPlayerHandler creates a new PlayerHandler.
func NewPlayerHandler(playerService *service.PlayerService) *PlayerHandler {
	return &PlayerHandler{playerService: playerService}
}

// SearchPlayer handles GET /api/v1/players/search?gamertag=&platform=
func (h *PlayerHandler) SearchPlayer(w http.ResponseWriter, r *http.Request) {
	gamertag := r.URL.Query().Get("gamertag")
	platform := r.URL.Query().Get("platform")

	if gamertag == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(apiError{
			Error:   "invalid_request",
			Message: "gamertag query parameter is required",
		})
		return
	}

	if platform == "" {
		platform = "uno"
	}

	title := r.URL.Query().Get("title")
	mode := r.URL.Query().Get("mode")

	result, err := h.playerService.SearchPlayer(r.Context(), platform, gamertag, title, mode)
	if err != nil {
		writeAPIError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// GetStats handles GET /api/v1/players/{platform}/{gamertag}/stats?mode=
func (h *PlayerHandler) GetStats(w http.ResponseWriter, r *http.Request) {
	platform := chi.URLParam(r, "platform")
	gamertag := chi.URLParam(r, "gamertag")
	title := r.URL.Query().Get("title")
	mode := r.URL.Query().Get("mode")

	stats, err := h.playerService.GetPlayerStats(r.Context(), platform, gamertag, title, mode)
	if err != nil {
		writeAPIError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}
