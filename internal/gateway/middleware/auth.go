package middleware

import (
	"errors"
	"github.com/cynxees/janus-gateway/internal/context"
	"github.com/cynxees/janus-gateway/internal/dependencies/config"
	"github.com/cynxees/janus-gateway/internal/dependencies/logger"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
	UserId   uint64 `json:"user_id"`
}

func GenerateToken(username string, userId uint64) (string, error) {
	claims := &Claims{
		Username: username,
		UserId:   userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Using default 24h since config uses string format
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Config.JWT.Secret))
}

func PublicAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				// No token, proceed without auth
				next.ServeHTTP(w, r)
				return
			}
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		tokenStr := cookie.Value
		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Config.JWT.Secret), nil
		})

		if err != nil || !token.Valid {
			// Invalid, continue without auth
			next.ServeHTTP(w, r.WithContext(r.Context()))
			return
		}

		// Add username to context
		ctx := context.SetKey(r.Context(), context.KeyUsername, claims.Username)
		ctx = context.SetUserId(ctx, claims.UserId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func PrivateAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				// No Token - Unauthorized
				http.Error(w, "Unauthorized, No Token in cookie", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		tokenStr := cookie.Value
		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Config.JWT.Secret), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Add username to context
		ctx := context.SetKey(r.Context(), context.KeyUsername, claims.Username)
		ctx = context.SetUserId(ctx, claims.UserId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		logger.Debug("CORS Middleware: Processing request")

		if !config.Config.CORS.Enabled {
			next.ServeHTTP(w, r)
			return
		}

		origin := r.Header.Get("Origin")
		logger.Debug("CORS Middleware: Origin: " + origin)

		if origin != "" {
			for _, allowedOrigin := range config.Config.CORS.Origins {
				logger.Debug("CORS Middleware: Checking allowed origin: " + allowedOrigin)
				if origin == allowedOrigin {
					logger.Debug("CORS Middleware: Origin allowed: " + origin)
					w.Header().Set("Access-Control-Allow-Origin", origin)
					w.Header().Add("Vary", "Origin")
					w.Header().Set("Access-Control-Allow-Credentials", "true")
					break
				}
			}
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
