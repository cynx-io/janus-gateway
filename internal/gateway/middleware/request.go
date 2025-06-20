package middleware

import (
	"bytes"
	"encoding/json"
	pb "github.com/cynxees/cynx-core/proto/gen"
	"github.com/cynxees/cynx-core/src/context"
	"github.com/cynxees/cynx-core/src/logger"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"log"
	"net/http"
	"time"
)

func BaseRequestHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger.Debug(ctx, "[BASE REQ] Processing request")

		reqId := uuid.New().String()
		userId := context.GetUserId(ctx)
		username := context.GetKey(ctx, context.KeyUsername)

		origin := r.Header.Get("Origin") // e.g. https://example.com
		timestamp := time.Now()

		baseReq := &pb.BaseRequest{
			RequestId:     reqId,
			RequestOrigin: origin,
			RequestPath:   r.URL.Path,
			UserId:        userId,
			Username:      username,
		}
		ctx, err := context.SetBaseRequest(ctx, baseReq)

		// Inject base request into body if JSON
		if baseReq != nil && r.Header.Get("Content-Type") == "application/json" {
			var bodyMap map[string]interface{}

			bodyBytes, err := io.ReadAll(r.Body)
			if err != nil {
				log.Printf("Failed to read request body: %v", err)
				http.Error(w, "Invalid request body", http.StatusBadRequest)
				return
			}

			if len(bodyBytes) > 0 {
				if err := json.Unmarshal(bodyBytes, &bodyMap); err != nil {
					log.Printf("Invalid JSON in request body: %v", err)
					http.Error(w, "Invalid JSON", http.StatusBadRequest)
					return
				}

				// Inject baseRequest
				baseMap := map[string]interface{}{
					"request_id":     baseReq.RequestId,
					"request_origin": baseReq.RequestOrigin,
					"request_path":   baseReq.RequestPath,
					"timestamp":      timestamppb.New(timestamp),
					"user_id":        baseReq.UserId,
					"username":       baseReq.Username,
				}
				bodyMap["base"] = baseMap

				// Replace the body
				updatedBody, err := json.Marshal(bodyMap)
				if err != nil {
					log.Printf("Failed to marshal updated body: %v", err)
					http.Error(w, "Internal Error", http.StatusInternalServerError)
					return
				}
				r.Body = io.NopCloser(bytes.NewBuffer(updatedBody))
			}
		}

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("Failed to set base request in context: %v", err)
			return
		}

		logger.Debug(ctx, "[BASE REQ] Success set for: ", reqId, " (UserID: ", userId, ", Username: ", username, ")")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
