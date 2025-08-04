package janus

import (
	"encoding/json"
	"github.com/cynx-io/janus-gateway/internal/dependencies/config"
	"github.com/cynx-io/janus-gateway/internal/helper"
	"github.com/cynx-io/janus-gateway/internal/session"
	"net/http"
	"net/url"
)

func (h *GatewayHandler) Auth0Logout(w http.ResponseWriter, r *http.Request) {

	siteKey, err := helper.GetSiteKey(r)
	if err != nil {
		http.Error(w, "Failed to get site key: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Clear the session
	err = session.ClearSession(w, r)
	if err != nil {
		http.Error(w, "Failed to clear session: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Build Auth0 logout URL
	logoutURL := "https://" + config.Config.Auth0.Domain + "v2/logout"
	params := url.Values{}
	params.Add("client_id", config.Config.Sites.Get(siteKey).Auth0.ClientId)
	if config.Config.Sites.Get(siteKey).Auth0.FrontendUrl != "" {
		params.Add("returnTo", config.Config.Sites.Get(siteKey).Auth0.FrontendUrl)
	}

	// Check if client wants to redirect to Auth0 logout
	if r.URL.Query().Get("type") == "full" || r.URL.Query().Get("redirect") == "true" {
		fullLogoutURL := logoutURL + "?" + params.Encode()
		http.Redirect(w, r, fullLogoutURL, http.StatusTemporaryRedirect)
		return
	}

	// Return JSON response with logout URL
	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"success":    true,
		"message":    "Logged out successfully",
		"logout_url": logoutURL + "?" + params.Encode(),
	}

	json.NewEncoder(w).Encode(response)
}
