package middleware

import (
	"bytes"
	"encoding/json"
	contextcore "github.com/cynxees/cynx-core/src/context"
	"github.com/cynxees/cynx-core/src/logger"

	"io"
	"log"
	"net/http"
	"time"
)

func LogRequestHandler(next http.Handler) http.Handler {
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
		}

		go func() {

			referer := r.Header.Get("Referer")      // e.g. https://example.com/page
			host := r.Host                          // e.g. api.myservice.com
			userAgent := r.Header.Get("User-Agent") // browser or bot details

			baseReq := contextcore.GetBaseRequest(ctx)

			var bodyMap map[string]interface{}
			_ = json.Unmarshal(bodyBytes, &bodyMap)

			if base, ok := bodyMap["base"].(map[string]interface{}); ok {
				if ts, ok := base["timestamp"].(map[string]interface{}); ok {
					seconds, _ := ts["seconds"].(float64)
					nanos, _ := ts["nanos"].(float64)
					t := time.Unix(int64(seconds), int64(nanos))
					base["timestamp"] = t.Format(time.RFC3339Nano)
				}
			}

			modifiedBodyBytes, _ := json.Marshal(bodyMap)

			logEntry := logger.TrxEntry{
				Timestamp:     time.Now(),
				UserId:        baseReq.UserId,
				Username:      baseReq.Username,
				RequestId:     baseReq.RequestId,
				RequestOrigin: baseReq.RequestOrigin,
				IpAddress:     baseReq.IpAddress,
				Endpoint:      r.Method + " " + r.URL.Path,
				Host:          host,
				Referer:       referer,
				UserAgent:     userAgent,
				Type:          "REQUEST",
				Body:          json.RawMessage(modifiedBodyBytes),
			}

			// Log to Elastic (ignore error or handle it)
			if err := logger.LogTrxElasticsearch(ctx, logEntry); err != nil {
				logger.Error(ctx, "Failed to log to ES: ", err)
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

func LogResponseHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// wrap the writer
		cw := newCaptureWriter(w)

		// let the handler run and write to our captureWriter
		next.ServeHTTP(cw, r)
		finishedAt := time.Now()

		go func() {
			ctx := r.Context()
			baseReq := contextcore.GetBaseRequest(ctx)

			referer := r.Header.Get("Referer")      // e.g. https://example.com/page
			host := r.Host                          // e.g. api.myservice.com
			userAgent := r.Header.Get("User-Agent") // browser or bot details

			// log to Elastic
			entry := logger.TrxEntry{
				Timestamp:     finishedAt,
				UserId:        baseReq.UserId,
				Username:      baseReq.Username,
				RequestId:     baseReq.RequestId,
				RequestOrigin: baseReq.RequestOrigin,
				IpAddress:     baseReq.IpAddress,
				Endpoint:      r.Method + " " + r.URL.Path,
				Host:          host,
				Referer:       referer,
				UserAgent:     userAgent,
				Type:          "RESPONSE",
				Body:          json.RawMessage(cw.body.Bytes()),
			}
			if err := logger.LogTrxElasticsearch(ctx, entry); err != nil {
				logger.Error(ctx, "response logging failed", err.Error())
			}
		}()
	})
}
