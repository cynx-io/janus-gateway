package middleware

import (
	pb "github.com/cynxees/janus-gateway/api/proto/gen/core"
	"github.com/cynxees/janus-gateway/internal/context"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net/http"
	"time"
)

func BaseRequestHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		reqId := uuid.New().String()
		userId := context.GetUserId(ctx)
		username := context.GetKey(ctx, context.KeyUsername)

		origin := r.Header.Get("Origin") // e.g. https://example.com
		timestamp := time.Now()

		ctx, err := context.SetBaseRequest(ctx, &pb.BaseRequest{
			RequestId:     reqId,
			RequestOrigin: origin,
			RequestPath:   r.URL.Path,
			Timestamp:     timestamppb.New(timestamp),
			UserId:        userId,
			Username:      username,
		})
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("Failed to set base request in context: %v", err)
			return
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
