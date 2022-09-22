package middleware

import (
	"context"
	"net/http"

	"iseage/bank/internal/usecase"
)

func TranslationCtx(tuc *usecase.TranslationUseCase) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(context.WithValue(r.Context(), "usecase.translation", tuc))
			next.ServeHTTP(w, r)
		})
	}
}
