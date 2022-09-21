// Package v1 implements routing paths. Each services in own file.
package www

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"

	"iseage/bank/pkg/logger"
)

func NewRouter(handler *chi.Mux, l logger.Interface) {
	// Options
	handler.Use(middleware.Logger)
	handler.Use(middleware.Recoverer)

	handler.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello world"))
	})
}
