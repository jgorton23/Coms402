// Package app configures and runs application.
package app

import (
	"fmt"
	"log"
	"os"

	"github.com/samber/do"
	"github.com/sirupsen/logrus"
	logrusadapter "logur.dev/adapter/logrus"
	"logur.dev/logur"

	"github.com/MatthewBehnke/exampleGoApi/internal/delivery/controller"
	"github.com/MatthewBehnke/exampleGoApi/internal/delivery/controller/http/api"
	"github.com/MatthewBehnke/exampleGoApi/internal/usecase"
	"github.com/MatthewBehnke/exampleGoApi/internal/usecase/repo"
	"github.com/MatthewBehnke/exampleGoApi/pkg/database"
)

// Run creates objects via constructors.
func Run() {
	l := NewLogger(LogConfig{
		Format:  "json", //logfmt or json
		Level:   "info",
		NoColor: true,
	})

	log.SetFlags(0) // Remove all flags
	// Override the global standard library logger to make sure everything uses our logger
	log.SetOutput(logur.NewLevelWriter(&l, logur.Info))

	injector := do.New()

	do.Provide(injector, repo.NewCleanEnvService)
	do.Provide(injector, usecase.NewConfigUseCase("./config.yml", l.WithSubsystem("config")))

	config := do.MustInvoke[*usecase.ConfigUseCase](injector)

	// Database repository
	db, err := database.NewClient(config.Get().PG.URL, database.MaxPoolSize(config.Get().PG.PoolMax))

	if err != nil {
		l.Error(fmt.Errorf("app - Run - postgres.New: %w", err).Error())
	}

	defer db.Client.Close()

	userRepo := repo.NewUserRepo(db)

	authBossUseCase := usecase.NewAuthBossUseCase(userRepo, l)

	httpV1UseCase := api.NewHttpV1(userRepo, l)

	if err != nil {
		l.Error(err.Error())
	}

	controller.New(config.Get(), l, *authBossUseCase, *httpV1UseCase)

	injector.Shutdown()
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
