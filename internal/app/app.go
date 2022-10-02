// Package app configures and runs application.
package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	chimiddleware "github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	logrusadapter "logur.dev/adapter/logrus"
	"logur.dev/logur"

	"iseage/bank/config"
	"iseage/bank/internal/delivery/controller/http/api"
	"iseage/bank/internal/delivery/controller/http/www"
	"iseage/bank/internal/delivery/middleware"
	"iseage/bank/internal/usecase"
	"iseage/bank/internal/usecase/repo"
	"iseage/bank/pkg/database"
	"iseage/bank/pkg/httpserver"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {

	l := NewLogger(LogConfig{
		Format:  "logfmt", //logfmt or json
		Level:   "info",
		NoColor: true,
	})

	// Override the global standard library logger to make sure everything uses our logger
	log.SetOutput(logur.NewLevelWriter(&l, logur.Info))

	// Database repository
	db, err := database.NewClient(cfg.PG.URL, database.MaxPoolSize(cfg.PG.PoolMax))

	if err != nil {
		l.Error(fmt.Errorf("app - Run - postgres.New: %w", err).Error())
	}

	defer db.Client.Close()

	userUseCase := usecase.NewUserUseCase(repo.NewUserRepo(db), l)

	if err != nil {
		l.Error(err.Error())
	}
	// HTTP Server
	handler := chi.NewRouter()
	handler.Use(chimiddleware.RequestID)
	handler.Use(chimiddleware.RealIP)
	handler.Use(chimiddleware.Recoverer)
	handler.Use(middleware.NewStructuredLogger(l))

	// loading usecase's onto context
	handler.Use(middleware.LoggerCtx(l))
	// handler.Use(middleware.TranslationCtx(translationUseCase))

	// Mounting the backend
	handler.Route("/api", api.NewRouter(cfg, handler, l, *userUseCase))

	// Mounting the frontend
	handler.Route("/", www.NewRouter)

	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err).Error())
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err).Error())
	}
}

// NewLogger creates a new logger.
func NewLogger(config LogConfig) usecase.LoggerAdapter {
	l := logrus.New()

	l.SetOutput(os.Stdout)

	switch config.Format {
	case "json":
		l.SetFormatter(&logrus.JSONFormatter{})
	case "logfmt":
		l.SetFormatter(&logrus.TextFormatter{
			DisableColors:             config.NoColor,
			EnvironmentOverrideColors: true,
		})
	}
	if level, err := logrus.ParseLevel(config.Level); err == nil {
		l.SetLevel(level)
	}

	return *usecase.NewLoggerAdapter(logrusadapter.New(l))
}

type LogConfig struct {
	// Format specifies the output log format.
	// Accepted values are: json, logfmt
	Format string

	// Level is the minimum log level that should appear on the output.
	Level string

	// NoColor makes sure that no log output gets colorized.
	NoColor bool
}
