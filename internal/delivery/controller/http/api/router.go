// Package api implements routing paths. Each services in own file.
package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	v1 "iseage/bank/internal/delivery/controller/http/api/v1"
)

func NewRouter(r chi.Router) {

	r.Get("/metrics", promhttp.Handler().(http.HandlerFunc))

	r.Route("/v1", v1.NewRouter)
}
