//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type (
	Compose mg.Namespace
	Linter  mg.Namespace
	Server  mg.Namespace
)

// Start docker compose development dependencies
func (Compose) Up() error {
	fmt.Println("Starting Dependencies...")

	err := sh.Run("docker-compose", "up", "--build", "-d", "postgres")

	return err
}

// Start docker compose development dependencies
func (Compose) Down() error {
	fmt.Println("Stopping Dependencies...")

	return sh.Run("docker-compose", "down", "--remove-orphans")
}

// Show docker compose development dependency logs
func (Compose) Logs() error {
	return sh.Run("docker-compose", "logs", "-f")
}

// Clean docker compose development environment
func (Compose) Clean() error {
	mg.Deps(Compose.Down)

	return sh.Run("docker", "volume", "rm", "com-s-402c_pg-data")
}

// Run all Linters
func (Linter) All() error {
	mg.SerialDeps(Linter.Gci)

	// mg.SerialDeps(Linter.Gofumpt)

	mg.SerialDeps(Linter.Golangci)

	return nil
}

// Run Gci Linter
func (Linter) Gci() error {
	fmt.Println("Running GCI linter")

	version, err := GetGithubLatestRelease("daixiang0", "gci")
	if err != nil {
		return err
	}

	fmt.Println("Version: " + version)

	ok, err := sh.Exec(nil, os.Stdout, os.Stdout, "go", "install", "github.com/daixiang0/gci@"+version)

	if !ok || err != nil {
		fmt.Println("Failed to install required dependency. Exiting...")
		return err
	}

	moduleName, err := GetGoModuleName()
	if err != nil {
		return err
	}

	return sh.Run("gci", "write", ".", "--skip-generated", "-s", "standard", "-s", "default", "-s", "blank", "-s", "\"prefix("+moduleName+")\"")
}

// // Run Gofumpt Linter
// func (Linter) Gofumpt() error {
// 	fmt.Println("Running gofumpt linter")

// 	version, err := GetGithubLatestRelease("mvdan", "gofumpt")
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Println("Version: " + version)

// 	ok, err := sh.Exec(nil, os.Stdout, os.Stdout, "go", "install", "mvdan.cc/gofumpt@"+version)

// 	if !ok || err != nil {
// 		fmt.Println("Failed to install required dependency. Exiting...")
// 		return err
// 	}

// 	return sh.Run("gofumpt", "-l", "-w", ".")
// }

// Run golangci-lint
func (Linter) Golangci() error {
	fmt.Println("Running Golangci linter")

	path, err := os.Getwd()
	if err != nil {
		return err
	}

	version, err := GetGithubLatestRelease("golangci", "golangci-lint")
	if err != nil {
		return err
	}

	fmt.Println("Version: " + version)

	ok, err := sh.Exec(nil, os.Stdout, os.Stdout, "docker", "run", "--rm", "-v", path+":/app", "-w", "/app", "golangci/golangci-lint:"+version, "golangci-lint", "run")

	if !ok {
		fmt.Println("Program failed to run")

		return nil
	}
	return err
}

// Start the Server
func (Server) Run() error {
	mg.Deps(Compose.Up)

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
