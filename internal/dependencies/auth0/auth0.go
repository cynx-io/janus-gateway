package auth0

import (
	"context"
	"encoding/gob"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/cynx-io/janus-gateway/internal/constant"
	"github.com/cynx-io/janus-gateway/internal/dependencies/config"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"net/http"
	"time"
)

var (
	Provider *oidc.Provider
	Verifier map[constant.SiteKey]*oidc.IDTokenVerifier
	Store    map[constant.SiteKey]*sessions.CookieStore
	Oauth2   map[constant.SiteKey]*oauth2.Config
)

func Init() {
	gob.Register(map[string]interface{}{})
	gob.Register(time.Time{})

	Verifier = make(map[constant.SiteKey]*oidc.IDTokenVerifier)
	Store = make(map[constant.SiteKey]*sessions.CookieStore)
	Oauth2 = make(map[constant.SiteKey]*oauth2.Config)

	ctx := context.Background()
	var err error

	Provider, err = oidc.NewProvider(ctx, "https://"+config.Config.Auth0.Domain)
	if err != nil {
		panic("Failed to create OIDC provider on " + config.Config.Auth0.Domain + ": " + err.Error())
	}

	config.Config.Sites.Iterate(func(key constant.SiteKey, cfg config.SiteConfig) {
		Oauth2[key] = &oauth2.Config{
			ClientID:     cfg.Auth0.ClientId,
			ClientSecret: cfg.Auth0.ClientSecret,
			RedirectURL:  cfg.Auth0.CallbackUrl,
			Endpoint:     Provider.Endpoint(),
			Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
		}

		oidcConfig := &oidc.Config{
			ClientID: cfg.Auth0.ClientId,
		}

		Verifier[key] = Provider.Verifier(oidcConfig)

		sessionSecret := cfg.Auth0.SessionSecret
		Store[key] = sessions.NewCookieStore(
			[]byte(sessionSecret),      // hash key (must be 32 or 64 bytes)
			[]byte(sessionSecret[:32]), // encryption key (must be 32 bytes for AES-256)
		)
		Store[key].Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
			Domain:   cfg.Domain,
		}
	})

}
