//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Test mg.Namespace

// This was a fun time getting e2e testing working correct....
// Some useful sites that I found information at that helped me get
// The coverage system working in vscode
//
// Converting the golang cover .out to the Cobertura Format (more useful and better support)
// https://stackoverflow.com/a/31440740
// https://github.com/Shippable/support/issues/3178#issuecomment-261732326
//
// Getting inline coverage reports in vscode. (coverage-gutters supports the Cobertura Format)
// https://marketplace.visualstudio.com/items?itemName=ryanluker.vscode-coverage-gutters
//
// https://go.dev/blog/integration-test-coverage

// Generate e2e coverage report
func (t Test) GenCoverageE2E() error {
	// mg.Deps(Compose.UpPostgres)

	sh.Rm("./cov.out")

	fmt.Println("Generating test binary")

	ok, err := sh.Exec(nil, os.Stdout, os.Stderr, "go", "test", "-c", "./internal/app", "-cover", "-covermode=set", "-coverpkg=./internal/...")

	if !ok && err != nil {
		fmt.Println("Program failed to run")
		fmt.Println(err)
		return nil
	}

	// Run the server in the background so we can hit it with out e2e tests
	go func() {
		fmt.Println("Starting test binary")
		ok, err = sh.Exec(nil, os.Stdout, os.Stdout, "./app.test", "-test.coverprofile", "cov.out")

		if !ok && err != nil {
			fmt.Println("Program failed to run")
			fmt.Println(err)
		}
	}()

	// TODO use a channel to verify the server is up first
	time.Sleep(10 * time.Second)

	_, _ = sh.Exec(nil, os.Stdout, os.Stderr, "go", "test", "-count=1", "./test/e2e/...")

	// this stops the server so we can get a coverage report
	_, _ = sh.Exec(nil, os.Stdout, os.Stderr, "curl", "localhost:19999")
	sh.Rm("./app.test")

	fmt.Println("Coverage report generated and automatically updated inline coverage report.")
	fmt.Println("You can now run test:CoverageReportWeb to view the report in the browser")

	mg.Deps(Test.CoverageReportInline, Test.CoverageReport)

	return err
}

func (Test) CoverageReport() error {
	//TODO get install package https://github.com/axw/gocov

	cmd := "gocov convert cov.out | gocov report > cov.txt"
	_, err := sh.Exec(nil, os.Stdout, os.Stderr, "bash", "-c", cmd)

	return err
}

func (Test) CoverageReportInline() error {
	//TODO get install package github.com/AlekSi/gocov-xml, https://github.com/axw/gocov

	cmd := "gocov convert cov.out | gocov-xml > cov.xml"
	_, err := sh.Exec(nil, os.Stdout, os.Stderr, "bash", "-c", cmd)

	return err
}

func (Test) CoverageReportWeb() error {
	//TODO get install package github.com/matm/gocov-html, https://github.com/axw/gocov

	cmd := "gocov convert cov.out | gocov-html -t kit > cov.html"
	_, err := sh.Exec(nil, os.Stdout, os.Stderr, "bash", "-c", cmd)

	cmd = "open cov.html"
	_, err = sh.Exec(nil, os.Stdout, os.Stderr, "bash", "-c", cmd)

	return err
}
