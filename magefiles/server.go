//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Server mg.Namespace

// Start the Server
func (Server) Run() error {
	mg.Deps(Compose.UpDev)

	fmt.Println("Running Server")

	env := map[string]string{
		"CGO_ENABLED": "0",
		"PG_URL":      "postgres://user:pass@localhost:5432/postgres?sslmode=disable",
		"APP_HOST":    "localhost:8082",
	}

	ok, err := sh.Exec(env, os.Stdout, os.Stderr, "go", "run", "./cmd/app/main.go")

	if !ok {
		fmt.Println("Program failed to run")

		return nil
	}
	return err
}
