package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/samber/do"

	v1 "git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/controller/http/v1"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/usecase"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/usecase/repo"
)

func main() {
	injector := do.New()

	config := repo.NewConfigRepo()
	err := config.Load("config.yml")

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
	// do.Provide(injector, repo.NewGorillaSessionRepo) // Gorilla has been deprecated
	do.Provide(injector, repo.NewSCSRepo)
	do.Provide(injector, repo.NewSCSSessionRepo)
	do.Provide(injector, repo.NewUserToCompanyRepo)
	do.Provide(injector, repo.NewCompanyRepo)

	do.Provide(injector, usecase.NewLogger)
	do.Provide(injector, usecase.NewAuthBossLogger)
	do.Provide(injector, usecase.NewAuthBossServer)
	do.Provide(injector, usecase.NewAuthbossSession)
	do.Provide(injector, usecase.NewHTTPAuthorization)
	do.Provide(injector, usecase.NewCompany)
	do.Provide(injector, usecase.NewUserToCompany)
	do.Provide(injector, usecase.NewUser)

	// HTTP stuff
	do.Provide(injector, v1.NewHttpAuthenticator)
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
