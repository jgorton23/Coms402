// Package app configures and runs application.
package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/samber/do"
	logrusadapter "logur.dev/adapter/logrus"
	"logur.dev/logur"

	"github.com/MatthewBehnke/exampleGoApi/internal/delivery/controller"
	"github.com/MatthewBehnke/exampleGoApi/internal/delivery/controller/http/api"
	"github.com/MatthewBehnke/exampleGoApi/internal/usecase"
	"github.com/MatthewBehnke/exampleGoApi/internal/usecase/repo"
)

// Run creates objects via constructors.
func Run() {
	injector := do.New()

	config := repo.NewCleanEnvService()
	config.Load("./config.yml")
	do.ProvideValue(injector, config.Get())

	do.Provide(injector, repo.NewLogrusService())
	logrus := do.MustInvoke[logrusadapter.Logger](injector)

	do.Provide(injector, usecase.NewLoggerUseCase(&logrus))
	l := do.MustInvoke[*usecase.LoggerUseCase](injector)

	// Remove all flags
	log.SetFlags(0)
	// Override the global standard library logger to make sure everything uses our logger
	log.SetOutput(logur.NewLevelWriter(l, logur.Info))

	do.Provide(injector, repo.NewDatabaseService)
	do.Provide(injector, repo.NewDataBaseServiceUser)
	do.Provide(injector, usecase.NewAuthBossUseCase)
	do.Provide(injector, api.NewHttpV1)
	do.Provide(injector, controller.New)

	controller := do.MustInvoke[*controller.Controller](injector)

	controller.Run()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Print("app - Run - signal: " + s.String())
	}

	// Shutdown
	err := injector.Shutdown()
	if err != nil {
		log.Print(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err).Error())
	}
}
