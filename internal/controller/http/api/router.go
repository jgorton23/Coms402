// Package api implements routing paths. Each services in own file.
package api

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	v1 "iseage/bank/internal/controller/http/api/v1"
	"iseage/bank/internal/usecase"
	"iseage/bank/pkg/logger"
)

func NewRouter(handler *chi.Mux, l logger.Interface, t usecase.Translation) {
	// Options

	handler.Route("/api", func(r chi.Router) {

		r.Use(middleware.Logger)
		r.Use(middleware.Recoverer)

		// r.Use(middleware.Heartbeat("/healthz"))

		r.Get("/metrics", promhttp.Handler().(http.HandlerFunc))

		v1.NewRouter(r.(*chi.Mux), l, t)

	})
}
