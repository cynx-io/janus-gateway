package middleware

import (
	"context"
	"fmt"
	contextcore "github.com/cynx-io/cynx-core/src/context"
	"github.com/cynx-io/cynx-core/src/logger"
	"github.com/cynx-io/cynx-core/src/types/usertype"
	"github.com/cynx-io/janus-gateway/internal/dependencies/auth0"
	"github.com/cynx-io/janus-gateway/internal/dependencies/config"
	"github.com/cynx-io/janus-gateway/internal/helper"
	"github.com/cynx-io/janus-gateway/internal/session"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/oauth2"
	"net/http"
	"strconv"
	"time"
)

type Claims struct {
	jwt.RegisteredClaims
	Username string            `json:"username"`
	UserId   int32             `json:"user_id"`
	UserType usertype.UserType `json:"user_type"`
}

func refreshToken(w http.ResponseWriter, r *http.Request, userSession *session.UserSession) error {
	if userSession.RefreshToken == "" {
		return &oauth2.RetrieveError{Response: &http.Response{StatusCode: 401}, Body: []byte("no refresh token")}
	}

	siteKey, _ := helper.GetSiteKey(r)
	tokenSource := auth0.Oauth2[siteKey].TokenSource(context.Background(), &oauth2.Token{
		RefreshToken: userSession.RefreshToken,
	})

	newToken, err := tokenSource.Token()
	if err != nil {
		return err
	}

	userSession.AccessToken = newToken.AccessToken
	if newToken.RefreshToken != "" {
		userSession.RefreshToken = newToken.RefreshToken
	}
	userSession.ExpiresAt = newToken.Expiry

	return session.SetSession(w, r, userSession)
}

func PublicAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger.Debug(ctx, "[PUBLIC AUTH] Processing request")

		userSession, err := session.GetSession(r)
		if err != nil || !userSession.Authenticated {
			// No session, proceed without auth
			next.ServeHTTP(w, r)
			return
		}

		// Check if token needs refresh (5 min buffer)
		if time.Now().Add(5 * time.Minute).After(userSession.ExpiresAt) {
			if refreshErr := refreshToken(w, r, userSession); refreshErr != nil {
				// Refresh failed, continue without auth
				logger.Error(ctx, "[PUBLIC AUTH] Token refresh failed: "+refreshErr.Error())
				next.ServeHTTP(w, r)
				return
			}
		}

		// Add user details to ctx (convert string UserID to int32)
		ctx = contextcore.SetKey(ctx, contextcore.KeyUsername, userSession.Name)
		ctx = contextcore.SetUserId(ctx, userSession.UserID)
		ctx = contextcore.SetUserType(ctx, 1) // Default user type

		logger.Debug(ctx, "[PUBLIC AUTH] Success set for: "+userSession.Name+" (UserID: "+string(userSession.UserID)+")")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func PrivateAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger.Debug(ctx, "[PRIVATE AUTH] Processing request")

		if r.Header.Get("X-Bypass-Auth") == "true" {
			ctx = bypassAuth(ctx, w, r)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		userSession, err := session.GetSession(r)
		if err != nil || !userSession.Authenticated {
			logger.Error(ctx, "[PRIVATE AUTH] No valid session")
			http.Error(w, "Unauthorized, No valid session", http.StatusUnauthorized)
			return
		}

		// Check if token needs refresh (5 min buffer)
		if time.Now().Add(5 * time.Minute).After(userSession.ExpiresAt) {
			if refreshErr := refreshToken(w, r, userSession); refreshErr != nil {
				logger.Error(ctx, "[PRIVATE AUTH] Token refresh failed: "+refreshErr.Error())
				if clearErr := session.ClearSession(w, r); clearErr != nil {
					logger.Error(ctx, "[PRIVATE AUTH] Failed to clear session: "+clearErr.Error())
				}
				http.Error(w, "Unauthorized, token refresh failed", http.StatusUnauthorized)
				return
			}
		}

		// Add user details to ctx (convert string UserID to int32)
		ctx = contextcore.SetKey(ctx, contextcore.KeyUsername, userSession.Name)
		ctx = contextcore.SetUserId(ctx, userSession.UserID)
		ctx = contextcore.SetUserType(ctx, 1)

		logger.Debug(ctx, "[PRIVATE AUTH] Success set for: "+userSession.Name+" (UserID: "+string(userSession.UserID)+")")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func bypassAuth(ctx context.Context, w http.ResponseWriter, r *http.Request) context.Context {

	if !config.Config.Bypass.Auth {
		logger.Error(ctx, "[PRIVATE AUTH] Bypass auth is disabled in config")
		http.Error(w, "Bypass auth is disabled", http.StatusForbidden)
		return ctx
	}

	// Get user details from headers
	userIDStr := r.Header.Get("X-Bypass-UserID")
	username := r.Header.Get("X-Bypass-Username")
	userTypeStr := r.Header.Get("X-Bypass-UserType")

	// Parse and validate UserID
	userID := int32(1) // default
	if userIDStr != "" {
		if id, err := strconv.ParseInt(userIDStr, 10, 32); err == nil {
			userID = int32(id)
		} else {
			logger.Error(ctx, "[PRIVATE AUTH] Invalid UserID format")
			http.Error(w, "Invalid UserID format", http.StatusBadRequest)
			return ctx
		}
	}

	// Set default username if not provided
	if username == "" {
		username = "Bypass User " + strconv.FormatInt(int64(userID), 10)
	}

	// Parse UserType
	userType := int32(1) // default
	if userTypeStr != "" {
		if ut, err := strconv.ParseInt(userTypeStr, 10, 32); err == nil {
			userType = int32(ut)
		}
	}

	// Set context values
	ctx = contextcore.SetKey(ctx, contextcore.KeyUsername, username)
	ctx = contextcore.SetUserId(ctx, userID)
	ctx = contextcore.SetUserType(ctx, userType)

	logger.Debug(ctx, fmt.Sprintf("[PRIVATE AUTH] Bypass success - User: %s, ID: %d, Type: %d",
		username, userID, userType))

	return ctx
}
