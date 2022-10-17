package middleware

import (
	"context"
	"net/http"

	"github.com/MatthewBehnke/exampleGoApi/internal/usecase"
)

func LoggerCtx(logger usecase.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "logger", logger)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}
