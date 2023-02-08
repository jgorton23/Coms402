package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/MatthewBehnke/apis/internal/app/usecase"
)

type loggerFields map[string]interface{}

func NewStructuredLogger(logger *usecase.Logger) func(next http.Handler) http.Handler {
	return middleware.RequestLogger(&StructuredLogger{logger})
}

type StructuredLogger struct {
	logger *usecase.Logger
}

func (l *StructuredLogger) NewLogEntry(r *http.Request) middleware.LogEntry {
	entry := &Entry{logger: l.logger}
	logFields := loggerFields{}

	if reqID := middleware.GetReqID(r.Context()); reqID != "" {
		logFields["req_id"] = reqID
	}

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	logFields["http_scheme"] = scheme
	logFields["http_proto"] = r.Proto
	logFields["http_method"] = r.Method

	logFields["remote_addr"] = r.RemoteAddr
	logFields["user_agent"] = r.UserAgent()

	logFields["uri"] = fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI)

	entry.logger.WithFields(logFields).Info("request started")

	return entry
}

type Entry struct {
	logger *usecase.Logger
}

func (e *Entry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
	e.logger.WithFields(loggerFields{
		"resp_status":       status,
		"resp_bytes_length": bytes,
		"resp_elapsed_ms":   float64(elapsed.Nanoseconds()) / 1000000.0,
	}).Info("request complete") // RequestLogger returns a logger handler using a custom LogFormatter.
}

func (e *Entry) Panic(v interface{}, stack []byte) {
	e.logger.WithFields(loggerFields{
		"stack": string(stack),
		"panic": fmt.Sprintf("%+v", v),
	})
}

// Helper methods used by the application to get the request-scoped
// logger entry and set additional fields between handlers.
//
// This is a useful pattern to use to set state on the entry as it
// passes through the handler chain, which at any point can be logged
// with a call to .Print(), .Info(), etc.

func GetLogEntry(r *http.Request) *usecase.Logger {
	entry := middleware.GetLogEntry(r).(*Entry)

	return entry.logger.WithFields(loggerFields{
		"req_id": middleware.GetReqID(r.Context()),
	})
}

// func LogEntrySetField(r *http.Request, key string, value interface{}) {
// 	if entry, ok := r.Context().Value(middleware.LogEntryCtxKey).(*Entry); ok {
// 		entry.logger.WithFields(loggerFields{
// 			key: value,
// 		})
// 	}
// }

// func LogEntrySetFields(r *http.Request, fields map[string]interface{}) {
// 	if entry, ok := r.Context().Value(middleware.LogEntryCtxKey).(*Entry); ok {
// 		entry.logger.WithFields(fields)
// 	}
// }
