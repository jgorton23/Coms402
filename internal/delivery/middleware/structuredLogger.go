package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/MatthewBehnke/exampleGoApi/internal/usecase"
)

func NewStructuredLogger(logger usecase.LoggerAdapter) func(next http.Handler) http.Handler {
	return middleware.RequestLogger(&StructuredLogger{logger})
}

type StructuredLogger struct {
	Logger usecase.LoggerAdapter
}

func (l *StructuredLogger) NewLogEntry(r *http.Request) middleware.LogEntry {
	entry := &Entry{Logger: *usecase.NewLoggerAdapter(&l.Logger)}
	logFields := usecase.LoggerFields{}

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
	// logFields["user_agent"] = r.UserAgent()

	logFields["uri"] = fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI)

	entry.Logger.WithFields(logFields).Info("request started")

	return entry
}

type Entry struct {
	Logger usecase.LoggerAdapter
}

func (e *Entry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
	e.Logger.WithFields(usecase.LoggerFields{
		"resp_status":       status,
		"resp_bytes_length": bytes,
		"resp_elapsed_ms":   float64(elapsed.Nanoseconds()) / 1000000.0,
	}).Info("request complete")
}

func (e *Entry) Panic(v interface{}, stack []byte) {
	e.Logger.WithFields(usecase.LoggerFields{
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

func GetLogEntry(r *http.Request) usecase.LoggerAdapter {
	entry := middleware.GetLogEntry(r).(*Entry)
	return entry.Logger
}

func LogEntrySetField(r *http.Request, key string, value interface{}) {
	if entry, ok := r.Context().Value(middleware.LogEntryCtxKey).(*Entry); ok {
		entry.Logger.WithFields(usecase.LoggerFields{
			key: value,
		})
	}
}

func LogEntrySetFields(r *http.Request, fields map[string]interface{}) {
	if entry, ok := r.Context().Value(middleware.LogEntryCtxKey).(*Entry); ok {
		entry.Logger.WithFields(fields)
	}
}
