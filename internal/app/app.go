// Package app configures and runs application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"iseage/bank/config"
	"iseage/bank/internal/controller/http/api"
	"iseage/bank/internal/controller/http/www"
	"iseage/bank/internal/usecase"
	"iseage/bank/internal/usecase/repo"
	"iseage/bank/internal/usecase/webapi"
	"iseage/bank/pkg/httpserver"
	"iseage/bank/pkg/logger"
	"iseage/bank/pkg/postgres"

	"github.com/go-chi/chi/v5"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// Use case
	translationUseCase := usecase.New(
		repo.New(pg),
		webapi.New(),
	)

	// HTTP Server
	handler := chi.NewRouter()
	www.NewRouter(handler, l)
	api.NewRouter(handler, l, translationUseCase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
