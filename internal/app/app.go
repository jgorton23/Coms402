// Package app configures and runs application.
package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/samber/do"
	"logur.dev/logur"

	"github.com/MatthewBehnke/exampleGoApi/internal/delivery/controller"
	"github.com/MatthewBehnke/exampleGoApi/internal/usecase"
	"github.com/MatthewBehnke/exampleGoApi/internal/usecase/repo"
)

func Run() {
	injector := do.New()

	config := repo.NewConfigRepo()
	err := config.Load("./config.yml")
	if err != nil {
		return
	}
	do.ProvideValue(injector, config.Get())

	do.Provide(injector, repo.NewLoggerRepo)
	do.Provide(injector, usecase.NewLogger)
	l := do.MustInvoke[*usecase.Logger](injector)

	// Remove all flags
	log.SetFlags(0)
	// Override the global standard library logger to make sure everything uses our logger
	log.SetOutput(logur.NewLevelWriter(l, logur.Info))

	do.Provide(injector, repo.NewDatabaseConnection)
	do.Provide(injector, repo.NewUserRepo)
	do.Provide(injector, repo.NewAuthorizationPolicyRepo)
	do.Provide(injector, repo.NewAuthorizationEnforcer)

	do.Provide(injector, usecase.NewAuthBossLogger)
	do.Provide(injector, usecase.NewAuthBossServer)
	do.Provide(injector, usecase.NewHttpAuthorization)

	//TODO Verify

	do.Provide(injector, controller.NewAuthenticator)
	do.Provide(injector, controller.NewHttpV1)
	do.Provide(injector, controller.New)

	c := do.MustInvoke[*controller.Controller](injector)

	c.Run()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Print("app - Run - signal: " + s.String())
	}

	// Shutdown
	err = injector.Shutdown()
	if err != nil {
		log.Print(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err).Error())
	}
}
