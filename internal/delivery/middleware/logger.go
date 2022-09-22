package middleware

import (
	"context"
	"net/http"

	"iseage/bank/internal/usecase"
)

func LoggerCtx(logger usecase.LoggerAdapter) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "logger", logger)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}
