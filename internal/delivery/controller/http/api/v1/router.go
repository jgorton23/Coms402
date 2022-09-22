// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"iseage/bank/internal/usecase"
)

func NewRouter(r chi.Router) {

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		logger, ok := r.Context().Value("logger").(usecase.LoggerAdapter)
		if !ok {
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		}

		logger.Warn("Errors")

		w.Write([]byte("Hello"))
	})
}
