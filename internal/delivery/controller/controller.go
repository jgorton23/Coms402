package controller

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	chimiddleware "github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/volatiletech/authboss/v3"
	_ "github.com/volatiletech/authboss/v3/auth"
	"github.com/volatiletech/authboss/v3/lock"
	_ "github.com/volatiletech/authboss/v3/logout"
	_ "github.com/volatiletech/authboss/v3/recover"
	_ "github.com/volatiletech/authboss/v3/register"
	"github.com/volatiletech/authboss/v3/remember"

	"iseage/bank/config"
	"iseage/bank/internal/delivery/controller/http/api"
	"iseage/bank/internal/delivery/middleware"
	"iseage/bank/internal/usecase"
	"iseage/bank/pkg/httpserver"
)

func New(cfg *config.Config, logger usecase.LoggerAdapter, abuc usecase.AuthBossUseCase, httpV1 api.HttpV1) {
	// HTTP Server
	mux := chi.NewRouter()
	mux.Use(chimiddleware.RequestID)
	mux.Use(chimiddleware.RealIP)
	mux.Use(chimiddleware.Recoverer)
	mux.Use(middleware.NewStructuredLogger(logger))

	// loading usecase's onto context
	mux.Use(middleware.LoggerCtx(logger))

	ab := newAuthentication(cfg, abuc, logger)

	mux.Use(ab.LoadClientStateMiddleware, remember.Middleware(ab))

	mux.Get("/metrics", promhttp.Handler().(http.HandlerFunc))

	// Mount the router to a path (this should be the same as the Mount path above)
	// mux in this example is a chi router, but it could be anything that can route to
	// the Core.Router.
	mux.Group(func(mux chi.Router) {
		mux.Use(authboss.ModuleListMiddleware(ab))
		mux.Mount("/auth", http.StripPrefix("/auth", ab.Config.Core.Router))
	})

	// Protected Routes
	mux.Group(func(r chi.Router) {
		r.Use(authboss.Middleware2(ab, authboss.RequireNone, authboss.RespondUnauthorized))
		r.Use(lock.Middleware(ab))

		// Openapi generated interfaces
		api.HandlerFromMuxWithBaseURL(httpV1, r, "/v1")
	})

	httpServer := httpserver.New(mux, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		logger.Info("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		logger.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err).Error())
	}

	// Shutdown
	err := httpServer.Shutdown()
	if err != nil {
		logger.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err).Error())
	}
}
