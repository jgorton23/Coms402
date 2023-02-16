//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Database mg.Namespace

// Generate Ent database bindings from schemas
func (Database) Generate() error {
	fmt.Println("Generating Ent Database Bindings")

	// env := map[string]string{
	// 	"CGO_ENABLED": "0",
	// 	"PG_URL":      "postgres://user:pass@localhost:5432/postgres?sslmode=disable",
	// 	"APP_HOST":    "localhost:8082",
	// }

	ok, err := sh.Exec(nil, os.Stdout, os.Stdout, "go", "run", "pkg/database/generate.go")

	if !ok {
		fmt.Println("Program failed to run")

		return nil
	}

	mg.SerialDeps(Linter.All)

	return err
}

// Create a new Ent Schema
func (Database) New(schema string) error {
	fmt.Println("Creating new schema")

	ok, err := sh.Exec(nil, os.Stdout, os.Stdout, "go", "run", "-mod=mod", "entgo.io/ent/cmd/ent", "init", "--target", "./pkg/database/models", schema)

	if !ok {
		fmt.Println("Program failed to run")

		return nil
	}

	mg.SerialDeps(Linter.All)

	return err
}

// Viz
func (Database) Viz() error {
	fmt.Println("View ER diagram")

	ok, err := sh.Exec(nil, os.Stdout, os.Stdout, "go", "run", "-mod=mod", "ariga.io/entviz", "./pkg/database/models")

	if !ok {
		fmt.Println("Program failed to run")

		return nil
	}

	return err
}
