package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/grovecj/warzone-stats-tracker/internal/service"
)

// MatchHandler holds dependencies for match endpoints.
type MatchHandler struct {
	matchService *service.MatchService
}

// NewMatchHandler creates a new MatchHandler.
func NewMatchHandler(matchService *service.MatchService) *MatchHandler {
	return &MatchHandler{matchService: matchService}
}

// GetMatches handles GET /api/v1/players/{platform}/{gamertag}/matches?limit=&offset=
func (h *MatchHandler) GetMatches(w http.ResponseWriter, r *http.Request) {
	platform := chi.URLParam(r, "platform")
	gamertag := chi.URLParam(r, "gamertag")

	limit := 20
	offset := 0
	if v := r.URL.Query().Get("limit"); v != "" {
		if parsed, err := strconv.Atoi(v); err == nil {
			limit = parsed
		}
	}
	if v := r.URL.Query().Get("offset"); v != "" {
		if parsed, err := strconv.Atoi(v); err == nil {
			offset = parsed
		}
	}

	title := r.URL.Query().Get("title")
	mode := r.URL.Query().Get("mode")

	result, err := h.matchService.GetRecentMatches(r.Context(), platform, gamertag, title, mode, limit, offset)
	if err != nil {
		writeAPIError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
