package middleware

import (
	"context"
	"github.com/cynx-io/cynx-core/src/logger"
	"github.com/cynx-io/janus-gateway/internal/constant"
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
		var siteKey constant.SiteKey

		// Check Origin header for CORS requests
		if origin != "" {
			logger.Debug(ctx, "CORS Middleware: Origin: "+origin)
			config.Config.Sites.Iterate(func(key constant.SiteKey, siteConfig config.SiteConfig) {
				logger.Debug(ctx, "CORS Middleware: Checking site: "+key)
				if allowedOrigin != "" {
					return
				}
				for _, o := range siteConfig.Urls {
					logger.Debug(ctx, "CORS Middleware: Checking url: "+o)
					if origin == o {
						allowedOrigin = origin
						siteKey = key
						break
					}
				}
			})
		} else {
			// Check host for direct API calls (no Origin header)
			host := "https://" + r.Host
			logger.Debug(ctx, "CORS Middleware: Host: "+host)
			config.Config.Sites.Iterate(func(key constant.SiteKey, siteConfig config.SiteConfig) {
				logger.Debug(ctx, "CORS Middleware: Checking API URL for site: "+key)
				if siteKey != "" {
					return
				}
				if host == siteConfig.ApiUrl {
					siteKey = key
				}
			})
		}

		logger.Debug(ctx, "[CORS] Allowed origin: "+allowedOrigin)
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

		// Site key is required - reject requests without valid origin
		if siteKey == "" {
			logger.Debug(ctx, "[CORS] No matching site found for origin: "+origin)
			w.WriteHeader(http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), constant.ContextKeySiteKey, siteKey)))
	})
}
