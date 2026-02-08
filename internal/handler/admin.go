package handler

import (
	"encoding/json"
	"net/http"

	"github.com/grovecj/warzone-stats-tracker/internal/codclient"
)

// AdminHandler holds dependencies for admin endpoints.
type AdminHandler struct {
	codClient codclient.CodClient
}

// NewAdminHandler creates a new AdminHandler.
func NewAdminHandler(codClient codclient.CodClient) *AdminHandler {
	return &AdminHandler{codClient: codClient}
}

type updateTokenRequest struct {
	Token string `json:"token"`
}

// UpdateToken handles POST /api/v1/admin/token to refresh the SSO token at runtime.
func (h *AdminHandler) UpdateToken(w http.ResponseWriter, r *http.Request) {
	var req updateTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "invalid_request",
			"message": "Request body must contain a JSON object with a 'token' field",
		})
		return
	}

	if req.Token == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "invalid_request",
			"message": "Token must not be empty",
		})
		return
	}

	h.codClient.UpdateToken(req.Token)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"message": "SSO token updated successfully",
	})
}
