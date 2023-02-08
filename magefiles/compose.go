//go:build mage
// +build mage

package main

import (
	"fmt"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Compose mg.Namespace

// Start docker compose development dependencies
func (Compose) UpDev() error {
	fmt.Println("Starting Dependencies...")

	err := sh.Run("docker-compose", "up", "--build", "-d", "postgres")

	return err
}

// Keeping for historical reference
// // Start docker compose development splunk
// func (Compose) UpSplunk() error {
// 	fmt.Println("Starting Splunk...")

// 	// Force docker to use AMD64 image because splunk does not offer a working arm image
// 	env := map[string]string{
// 		"DOCKER_DEFAULT_PLATFORM": "linux/amd64",
// 	}

// 	ok, err := sh.Exec(env, os.Stdout, os.Stderr, "docker-compose", "up", "--build", "-d", "so1")

// 	if !ok {
// 		fmt.Println("Program failed to run")

// 		return nil
// 	}
// 	return err
// }

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
