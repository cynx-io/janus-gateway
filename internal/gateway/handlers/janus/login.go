package janus

import (
	"github.com/cynx-io/cynx-core/src/helper/random"
	"github.com/cynx-io/janus-gateway/internal/dependencies/auth0"
	"github.com/cynx-io/janus-gateway/internal/helper"
	"github.com/cynx-io/janus-gateway/internal/session"
	"net/http"
)

func (h *GatewayHandler) Auth0Login(w http.ResponseWriter, r *http.Request) {
	state := random.RandomAlphanumerics(32, 32)

	siteKey, err := helper.GetSiteKey(r)
	if err != nil {
		http.Error(w, "Failed to get site key: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = session.SetState(w, r, state)
	if err != nil {
		http.Error(w, "Failed to save session: "+err.Error(), http.StatusInternalServerError)
		return
	}

	url := auth0.Oauth2[siteKey].AuthCodeURL(state)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
