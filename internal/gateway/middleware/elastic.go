package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	pb "github.com/cynxees/janus-gateway/api/proto/gen/core"
	"github.com/cynxees/janus-gateway/internal/context"
	"github.com/cynxees/janus-gateway/internal/dependencies/logger"
	"github.com/elastic/go-elasticsearch"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"log"
	"net/http"
	"time"
)

type LoggingMiddleware struct {
	ElasticClient *elasticsearch.Client
}

type LogEntry struct {
	Timestamp     time.Time `json:"timestamp"`
	UserId        *uint64   `json:"userId"`
	Username      *string   `json:"username"`
	RequestId     string    `json:"requestId"`
	RequestOrigin string    `json:"requestOrigin"`
	IpAddress     string    `json:"ipAddress"`
	Endpoint      string    `json:"endpoint"`
	Host          string    `json:"host"`
	Referer       string    `json:"referer,omitempty"`
	UserAgent     string    `json:"userAgent,omitempty"`
}

func logToElasticsearch(es *elasticsearch.Client, entry LogEntry) error {
	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	res, err := es.Index(
		"log-janus-gateway",
		bytes.NewReader(data),
		es.Index.WithDocumentType("_doc"),
	)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Error("Error closing response body: ", err)
		}
	}(res.Body)

	if res.IsError() {
		return fmt.Errorf("elasticsearch index error: %s", res.String())
	}
	return nil
}

func (m *LoggingMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		reqId := uuid.New().String()
		userId := context.GetUserId(ctx)
		username := context.GetKey(ctx, context.KeyUsername)

		referer := r.Header.Get("Referer")      // e.g. https://example.com/page
		origin := r.Header.Get("Origin")        // e.g. https://example.com
		host := r.Host                          // e.g. api.myservice.com
		userAgent := r.Header.Get("User-Agent") // browser or bot details
		timestamp := time.Now()

		logEntry := LogEntry{
			RequestId:     reqId,
			RequestOrigin: origin,
			IpAddress:     r.RemoteAddr,
			Endpoint:      r.Method + " " + r.URL.Path,
			Host:          host,
			Referer:       referer,
			UserAgent:     userAgent,
			UserId:        userId,
			Username:      username,
			Timestamp:     timestamp,
		}

		// Log to Elasticsearch (ignore error or handle it)
		if err := logToElasticsearch(m.ElasticClient, logEntry); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("Failed to log to ES: %v", err)
		}

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

		// pass request along
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
