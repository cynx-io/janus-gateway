package auth0

import (
	"context"
	"encoding/gob"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/cynx-io/janus-gateway/internal/dependencies/config"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"net/http"
	"time"
)

var (
	Provider *oidc.Provider
	Verifier *oidc.IDTokenVerifier
	Store    *sessions.CookieStore
	Oauth2   *oauth2.Config
)

func Init() {
	gob.Register(map[string]interface{}{})
	gob.Register(time.Time{})

	ctx := context.Background()
	var err error

	Provider, err = oidc.NewProvider(ctx, "https://"+config.Config.Auth0.Domain)
	if err != nil {
		panic("Failed to create OIDC provider on " + config.Config.Auth0.Domain + ": " + err.Error())
	}

	Oauth2 = &oauth2.Config{
		ClientID:     config.Config.Auth0.ClientId,
		ClientSecret: config.Config.Auth0.ClientSecret,
		RedirectURL:  config.Config.Auth0.CallbackUrl,
		Endpoint:     Provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	oidcConfig := &oidc.Config{
		ClientID: config.Config.Auth0.ClientId,
	}
	Verifier = Provider.Verifier(oidcConfig)

	sessionSecret := config.Config.Auth0.SessionSecret
	Store = sessions.NewCookieStore(
		[]byte(sessionSecret),      // hash key (must be 32 or 64 bytes)
		[]byte(sessionSecret[:32]), // encryption key (must be 32 bytes for AES-256)
	)
	Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7, // 7 days
		HttpOnly: true,
		Secure:   true, // Required for HTTPS
		SameSite: http.SameSiteLaxMode,
		Domain:   ".cynxio.com", // Allow cookies across subdomains
	}
}
