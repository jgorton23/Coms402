package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"

	"iseage/bank/pkg/database/ent"
)

const (
	_defaultMaxPoolSize = 1
)

// Database -.
type Database struct {
	maxPoolSize int

	Client *ent.Client
}

func NewClient(url string, opts ...Option) (*Database, error) {
	pg := &Database{
		maxPoolSize: _defaultMaxPoolSize,
	}

	// Custom options
	for _, opt := range opts {
		opt(pg)
	}

	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to: %v", err)
	}

	db.SetMaxOpenConns(pg.maxPoolSize)

	drv := entsql.OpenDB("postgres", db)

	pg.Client = ent.NewClient(ent.Driver(drv))

	// Run the auto migration tool.
	if err := pg.Client.Schema.Create(context.Background()); err != nil {
		return nil, fmt.Errorf("failed creating schema resources: %v", err)
	}
	log.Print("finished database setup")
	return pg, nil
}
