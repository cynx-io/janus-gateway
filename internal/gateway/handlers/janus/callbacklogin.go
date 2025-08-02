package janus

import (
	"context"
	"encoding/json"
	"github.com/cynx-io/janus-gateway/internal/dependencies/auth0"
	"github.com/cynx-io/janus-gateway/internal/dependencies/config"
	"github.com/cynx-io/janus-gateway/internal/session"
	"net/http"
)

func (h *GatewayHandler) Auth0CallbackLogin(w http.ResponseWriter, r *http.Request) {
	sessionState, err := session.GetState(r)
	if err != nil {
		http.Error(w, "Failed to get session", http.StatusInternalServerError)
		return
	}

	state := r.URL.Query().Get("state")
	if state != sessionState {
		http.Error(w, "Invalid state parameter", http.StatusBadRequest)
		return
	}

	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Code not found", http.StatusBadRequest)
		return
	}

	token, err := auth0.Oauth2.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "Failed to exchange code: "+err.Error(), http.StatusInternalServerError)
		return
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		http.Error(w, "No id_token field in token", http.StatusInternalServerError)
		return
	}

	idToken, err := auth0.Verifier.Verify(context.Background(), rawIDToken)
	if err != nil {
		http.Error(w, "Failed to verify ID Token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var claims map[string]interface{}
	if err := idToken.Claims(&claims); err != nil {
		http.Error(w, "Failed to get claims: "+err.Error(), http.StatusInternalServerError)
		return
	}

	userSession := &session.UserSession{
		UserID:        claims["sub"].(string),
		Email:         claims["email"].(string),
		Name:          claims["name"].(string),
		Authenticated: true,
	}

	err = session.SetSession(w, r, userSession)
	if err != nil {
		http.Error(w, "Failed to save session: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if config.Config.Auth0.FrontendUrl != "" {
		http.Redirect(w, r, config.Config.Auth0.FrontendUrl, http.StatusTemporaryRedirect)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"success": true,
		"user":    userSession,
		"message": "Login successful",
	}
	json.NewEncoder(w).Encode(response)
}
