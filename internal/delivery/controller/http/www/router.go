// Package v1 implements routing paths. Each services in own file.
package www

import (
	"iseage/bank/internal/usecase"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func NewRouter(handler *chi.Mux, l usecase.Logger) {
	// Options
	handler.Use(middleware.Recoverer)

	handler.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello world"))
	})
}
