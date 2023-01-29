//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Linter mg.Namespace

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

	return sh.Run("gci", "write", ".", "-s", "standard", "-s", "default", "-s", "blank", "-s", "\"prefix("+moduleName+")\"")
}

// Currently gofumpt conflicts with the wls linter so like.... Just keeping this code block here
// for future reference
//
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
