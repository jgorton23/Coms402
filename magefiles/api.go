//go:build mage
// +build mage

package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Api mg.Namespace

func (Api) Gen() error {
	fmt.Println("Starting Dependencies...")

	version, err := GetGithubLatestRelease("deepmap", "oapi-codegen")
	if err != nil {
		return err
	}

	fmt.Println("Version: " + version)

	ok, err := sh.Exec(nil, os.Stdout, os.Stdout, "go", "install", "github.com/deepmap/oapi-codegen/cmd/oapi-codegen@"+version)

	if !ok || err != nil {
		fmt.Println("Failed to install required dependency. Exiting...")
		return err
	}

	f, err := os.OpenFile("internal/app/controller/http/v1/apiv1.gen.go", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)

	if err != nil {
		return err
	}

	defer f.Close()

	w := bufio.NewWriter(f)

	ok, err = sh.Exec(nil, w, os.Stdout, "oapi-codegen", "--config", ".oapi-config.yml", "openapi.yml")

	if !ok {
		return err
	}

	return nil
}
