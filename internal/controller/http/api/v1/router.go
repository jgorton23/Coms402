// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"iseage/bank/internal/usecase"
	"iseage/bank/pkg/logger"
)

// NewRouter -.
// Swagger spec:
// @title       Go Clean Template API
// @description Using a translation service as an example
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(handler *chi.Mux, l logger.Interface, t usecase.Translation) {

	handler.Route("/v1", func(r chi.Router) {
		// newTranslationRoutes(r, t, l)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello"))
		})
	})
}
