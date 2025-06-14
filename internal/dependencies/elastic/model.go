package elastic

import (
	"encoding/json"
	"time"
)

type LogEntry struct {
	Timestamp     time.Time       `json:"timestamp"`
	UserId        *uint64         `json:"userId"`
	Username      *string         `json:"username"`
	RequestId     string          `json:"requestId"`
	RequestOrigin string          `json:"requestOrigin"`
	IpAddress     string          `json:"ipAddress"`
	Endpoint      string          `json:"endpoint"`
	Host          string          `json:"host"`
	Referer       string          `json:"referer,omitempty"`
	UserAgent     string          `json:"userAgent,omitempty"`
	Type          string          `json:"type,omitempty"` // e.g. "request", "response"
	Body          json.RawMessage `json:"body,omitempty"` // request or response body
}
