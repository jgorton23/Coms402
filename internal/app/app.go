// Package app configures and runs application.
package app

import (
	"fmt"
	v1 "github.com/MatthewBehnke/exampleGoApi/internal/delivery/http/v1"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/samber/do"

	"github.com/MatthewBehnke/exampleGoApi/internal/usecase"
	"github.com/MatthewBehnke/exampleGoApi/internal/usecase/repo"
)

func Run() {
	injector := do.New()

	config := repo.NewConfigRepo()
	err := config.Load("./config.yml")
	if err != nil {
		log.Fatal(err)
		return
	}

	do.ProvideValue(injector, config.Get())
	do.Provide(injector, repo.NewLoggerRepo)
	do.Provide(injector, repo.NewDatabaseConnection)
	do.Provide(injector, repo.NewUserRepo)
	do.Provide(injector, repo.NewAuthorizationPolicyRepo)
	do.Provide(injector, repo.NewAuthorizationEnforcer)
	do.Provide(injector, usecase.NewLogger)
	do.Provide(injector, usecase.NewAuthBossLogger)
	do.Provide(injector, usecase.NewAuthBossServer)
	do.Provide(injector, usecase.NewHttpAuthorization)

	//TODO Verify
	do.Provide(injector, usecase.NewHttpAuthenticator)
	do.Provide(injector, v1.NewHttpV1)
	do.Provide(injector, v1.NewHttpV1Router)

	c := do.MustInvoke[*v1.HttpV1Router](injector)
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
