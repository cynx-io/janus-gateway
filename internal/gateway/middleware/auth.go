package middleware

import (
	"errors"
	"github.com/cynxees/janus-gateway/internal/context"
	"github.com/cynxees/janus-gateway/internal/dependencies/config"
	"github.com/cynxees/janus-gateway/internal/dependencies/logger"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
	UserId   int32  `json:"user_id"`
}

func GenerateToken(username string, userId int32) (string, error) {
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
		logger.Debug("[PUBLIC AUTH] Processing request")
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
		logger.Debug("[PUBLIC AUTH] Success set for: " + claims.Username + " (UserID: " + strconv.Itoa(int(claims.UserId)) + ")")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func PrivateAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Debug("[PRIVATE AUTH] Processing request")

		cookie, err := r.Cookie("token")
		if err != nil {
			logger.Error("[PRIVATE AUTH] Error getting cookie: " + err.Error())
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
			logger.Error("[PRIVATE AUTH] Invalid token: " + err.Error())
			http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
			return
		}

		// Add username to context
		ctx := context.SetKey(r.Context(), context.KeyUsername, claims.Username)
		ctx = context.SetUserId(ctx, claims.UserId)

		logger.Debug("[PRIVATE AUTH] Success set for: ", claims.Username, " (UserID: ", claims.UserId, ")")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Debug("[CORS]: Processing request")

		if !config.Config.CORS.Enabled {
			next.ServeHTTP(w, r)
			return
		}

		origin := r.Header.Get("Origin")
		logger.Debug("CORS Middleware: Origin: " + origin)

		allowedOrigin := ""
		if origin != "" {
			for _, o := range config.Config.CORS.Origins {
				logger.Debug("CORS Middleware: Checking allowed origin: " + o)
				if origin == o {
					allowedOrigin = origin
					break
				}
			}
		}

		if allowedOrigin != "" {
			// Set only if origin is allowed, never '*'
			w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
			w.Header().Add("Vary", "Origin") // ensure caching varies by origin
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Connect-Protocol-Version")
		}

		// Handle preflight OPTIONS request early
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		logger.Debug("[CORS] Success set for origin: " + allowedOrigin)
		next.ServeHTTP(w, r)
	})
}
