package middleware

import (
	"context"
	contextcore "github.com/cynx-io/cynx-core/src/context"
	"github.com/cynx-io/cynx-core/src/logger"
	"github.com/cynx-io/cynx-core/src/types/usertype"
	"github.com/cynx-io/janus-gateway/internal/dependencies/auth0"
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
		userID, _ := strconv.ParseInt(userSession.UserID, 10, 32)
		ctx = contextcore.SetKey(ctx, contextcore.KeyUsername, userSession.Name)
		ctx = contextcore.SetUserId(ctx, int32(userID))
		ctx = contextcore.SetUserType(ctx, 1) // Default user type

		logger.Debug(ctx, "[PUBLIC AUTH] Success set for: "+userSession.Name+" (UserID: "+userSession.UserID+")")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func PrivateAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger.Debug(ctx, "[PRIVATE AUTH] Processing request")

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
		userID, _ := strconv.ParseInt(userSession.UserID, 10, 32)
		ctx = contextcore.SetKey(ctx, contextcore.KeyUsername, userSession.Name)
		ctx = contextcore.SetUserId(ctx, int32(userID))
		ctx = contextcore.SetUserType(ctx, 1) // Default user type

		logger.Debug(ctx, "[PRIVATE AUTH] Success set for: "+userSession.Name+" (UserID: "+userSession.UserID+")")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
