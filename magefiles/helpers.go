//go:build mage
// +build mage

package main

import (
	"context"
	"io/ioutil"

	"github.com/google/go-github/v50/github"
	"golang.org/x/mod/modfile"
)

func GetGithubLatestRelease(org, repo string) (string, error) {
	client := github.NewClient(nil)
	ctx := context.Background()

	latestRelease, _, err := client.Repositories.GetLatestRelease(ctx, org, repo)
	if err != nil {
		return "", err
	}

	return *latestRelease.TagName, nil
}

func GetGoModuleName() (string, error) {
	goModBytes, err := ioutil.ReadFile("go.mod")
	if err != nil {
		return "", err
	}

	modName := modfile.ModulePath(goModBytes)

	return modName, nil
}
