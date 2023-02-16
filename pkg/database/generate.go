//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	if err := entc.Generate("./pkg/database/models", &gen.Config{
		Package: "git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent",
		Target:  "./pkg/database/ent",
		Features: []gen.Feature{
			gen.FeatureLock,
			gen.FeatureUpsert,
		},
	}); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
