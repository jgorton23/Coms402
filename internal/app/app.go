// Package app configures and runs application.
package app

import (
	"fmt"
	"log"
	"os"

	"github.com/sirupsen/logrus"
	logrusadapter "logur.dev/adapter/logrus"
	"logur.dev/logur"

	"iseage/bank/config"
	"iseage/bank/internal/delivery/controller"
	"iseage/bank/internal/usecase"
	"iseage/bank/internal/usecase/repo"
	"iseage/bank/pkg/database"
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

	authBossUseCase := usecase.NewAuthBossUseCase(repo.NewUserRepo(db), l)

	if err != nil {
		l.Error(err.Error())
	}

	controller.New(cfg, l, *authBossUseCase)
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
