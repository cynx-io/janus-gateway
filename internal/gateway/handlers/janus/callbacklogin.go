package janus

import (
	"context"
	"encoding/json"
	gen "github.com/cynx-io/cynx-core/proto/gen"
	"github.com/cynx-io/cynx-core/src/logger"
	proto "github.com/cynx-io/janus-gateway/api/proto/gen/hermes"
	"github.com/cynx-io/janus-gateway/internal/dependencies/auth0"
	"github.com/cynx-io/janus-gateway/internal/dependencies/config"
	"github.com/cynx-io/janus-gateway/internal/helper"
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

	siteKey, err := helper.GetSiteKey(r)
	if err != nil {
		http.Error(w, "Failed to get site key: "+err.Error(), http.StatusInternalServerError)
		return
	}
	token, err := auth0.Oauth2[siteKey].Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "Failed to exchange code: "+err.Error(), http.StatusInternalServerError)
		return
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		http.Error(w, "No id_token field in token", http.StatusInternalServerError)
		return
	}

	idToken, err := auth0.Verifier[siteKey].Verify(context.Background(), rawIDToken)
	if err != nil {
		http.Error(w, "Failed to verify ID Token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var claims map[string]interface{}
	if err := idToken.Claims(&claims); err != nil {
		http.Error(w, "Failed to get claims: "+err.Error(), http.StatusInternalServerError)
		return
	}

	req := proto.UpsertUserRequest{
		Base: &gen.BaseRequest{},
	}
	req.Auth0Id = claims["sub"].(string)
	req.Email = claims["email"].(string)
	name := claims["name"].(string)
	req.Name = &name
	isActive := true
	req.IsActive = &isActive

	userResp, err := h.userClient.UpsertUser(r.Context(), &req)

	if err != nil {
		http.Error(w, "Failed to upsert user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if userResp == nil || userResp.User == nil {
		http.Error(w, "User not found in response", http.StatusInternalServerError)
		return
	}

	userSession := &session.UserSession{
		UserID:        string(userResp.User.Id),
		Email:         userResp.User.Email,
		Name:          userResp.User.Name,
		Authenticated: true,
		AccessToken:   token.AccessToken,
		RefreshToken:  token.RefreshToken,
		ExpiresAt:     token.Expiry,
	}

	err = session.SetSession(w, r, userSession)
	if err != nil {
		http.Error(w, "Failed to save session: "+err.Error(), http.StatusInternalServerError)
		return
	}

	redirectURL, err := session.GetRedirectURL(r)
	if err != nil || redirectURL == "" {
		siteKey, _ := helper.GetSiteKey(r)
		redirectURL = config.Config.Sites.Get(siteKey).Auth0.FrontendUrl
	}

	err = session.ClearRedirectURL(w, r)
	if err != nil {
		logger.Error(r.Context(), "Failed to clear redirect URL: ", err)
		return
	}

	if redirectURL != "" {
		http.Redirect(w, r, redirectURL, http.StatusTemporaryRedirect)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"success": true,
		"user":    userSession,
		"message": "Login successful",
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
