// Package v1 implements routing paths. Each services in own file.
package www

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter(r chi.Router) {

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

}