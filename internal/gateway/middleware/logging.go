package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/cynxees/janus-gateway/internal/context"
	"github.com/cynxees/janus-gateway/internal/dependencies/elastic"
	"github.com/cynxees/janus-gateway/internal/dependencies/logger"
	"io"
	"log"
	"net/http"
	"time"
)

type LoggingMiddleware struct {
	ElasticClient *elastic.Client
}

func (m *LoggingMiddleware) RequestHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		requestBody := r.Body
		var bodyBytes []byte

		if requestBody != nil && requestBody != http.NoBody {
			var err error
			bodyBytes, err = io.ReadAll(requestBody)
			if err != nil {
				log.Printf("Failed to read body: %v", err)
			}
			// Replace the body so the next handler can still read it
			r.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		} else {
			requestBody = http.NoBody
		}

		go func() {

			referer := r.Header.Get("Referer")      // e.g. https://example.com/page
			host := r.Host                          // e.g. api.myservice.com
			userAgent := r.Header.Get("User-Agent") // browser or bot details

			baseReq := context.GetBaseRequest(ctx)

			logEntry := elastic.LogEntry{
				Timestamp:     baseReq.Timestamp.AsTime(),
				UserId:        baseReq.UserId,
				Username:      baseReq.Username,
				RequestId:     baseReq.RequestId,
				RequestOrigin: baseReq.RequestOrigin,
				IpAddress:     r.RemoteAddr,
				Endpoint:      r.Method + " " + r.URL.Path,
				Host:          host,
				Referer:       referer,
				UserAgent:     userAgent,
				Type:          "REQUEST",
				Body:          json.RawMessage(bodyBytes),
			}

			// Log to Elasticsearch (ignore error or handle it)
			if err := m.ElasticClient.LogToElasticsearch(logEntry); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				log.Printf("Failed to log to ES: %v", err)
			}
		}()

		// pass request along
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type captureWriter struct {
	http.ResponseWriter
	body       *bytes.Buffer
	statusCode int
}

func newCaptureWriter(w http.ResponseWriter) *captureWriter {
	return &captureWriter{w, new(bytes.Buffer), http.StatusOK}
}

func (cw *captureWriter) WriteHeader(code int) {
	cw.statusCode = code
	cw.ResponseWriter.WriteHeader(code)
}

func (cw *captureWriter) Write(b []byte) (int, error) {
	cw.body.Write(b) // capture
	return cw.ResponseWriter.Write(b)
}

func (m *LoggingMiddleware) ResponseHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// wrap the writer
		cw := newCaptureWriter(w)

		// let the handler run and write to our captureWriter
		next.ServeHTTP(cw, r)

		go func() {
			ctx := r.Context()
			baseReq := context.GetBaseRequest(ctx)

			referer := r.Header.Get("Referer")      // e.g. https://example.com/page
			host := r.Host                          // e.g. api.myservice.com
			userAgent := r.Header.Get("User-Agent") // browser or bot details

			// log to Elasticsearch
			entry := elastic.LogEntry{
				Timestamp:     time.Now(),
				UserId:        baseReq.UserId,
				Username:      baseReq.Username,
				RequestId:     baseReq.RequestId,
				RequestOrigin: baseReq.RequestOrigin,
				IpAddress:     r.RemoteAddr,
				Endpoint:      r.Method + " " + r.URL.Path,
				Host:          host,
				Referer:       referer,
				UserAgent:     userAgent,
				Type:          "RESPONSE",
				Body:          json.RawMessage(cw.body.Bytes()),
			}
			if err := m.ElasticClient.LogToElasticsearch(entry); err != nil {
				logger.Error("response logging failed", err.Error())
			}
		}()
	})
}
