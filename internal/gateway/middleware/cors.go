package middleware

import (
	"github.com/cynx-io/cynx-core/src/logger"
	"github.com/cynx-io/janus-gateway/internal/dependencies/config"
	"net/http"
)

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger.Debug(ctx, "[CORS]: Processing request")

		if !config.Config.CORS.Enabled {
			next.ServeHTTP(w, r)
			return
		}

		origin := r.Header.Get("Origin")
		logger.Debug(ctx, "CORS Middleware: Origin: "+origin)

		allowedOrigin := ""
		if origin != "" {
			for _, o := range config.Config.CORS.Origins {
				logger.Debug(ctx, "CORS Middleware: Checking allowed origin: "+o)
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

		logger.Debug(ctx, "[CORS] Success set for origin: "+allowedOrigin)
		next.ServeHTTP(w, r)
	})
}
