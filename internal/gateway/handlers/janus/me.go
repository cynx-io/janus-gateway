package janus

import (
	"encoding/json"
	"github.com/cynx-io/janus-gateway/internal/session"
	"net/http"
)

func (h *GatewayHandler) Auth0Me(w http.ResponseWriter, r *http.Request) {
	userSession, err := session.GetSession(r)
	if err != nil || !userSession.Authenticated {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"user_id":       userSession.UserID,
		"email":         userSession.Email,
		"name":          userSession.Name,
		"authenticated": userSession.Authenticated,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
