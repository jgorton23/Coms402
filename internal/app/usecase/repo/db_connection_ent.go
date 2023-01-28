package repo

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/samber/do"

	"github.com/MatthewBehnke/apis/internal/app/domain"
	"github.com/MatthewBehnke/apis/pkg/database/ent"

	_ "github.com/lib/pq"
)

func NewDatabaseConnection(i *do.Injector) (*DatabaseConnection, error) {
	dbService := &DatabaseConnection{}

	config := do.MustInvoke[*domain.Config](i)

	db, err := sql.Open("postgres", config.PG.URL)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to: %w", err)
	}

	dbService.db = db

	dbService.db.SetMaxOpenConns(config.PG.PoolMax)

	drv := entsql.OpenDB("postgres", dbService.db)

	dbService.entClient = ent.NewClient(ent.Driver(drv))

	// Run the auto migration tool.
	if err := dbService.entClient.Schema.Create(context.Background()); err != nil {
		return nil, fmt.Errorf("failed creating schema resources: %w", err)
	}

	log.Print("database service started")

	return dbService, nil
}

type DatabaseConnection struct {
	db        *sql.DB
	entClient *ent.Client
}

func (s *DatabaseConnection) Client() *ent.Client {
	return s.entClient
}

func (s *DatabaseConnection) HealthCheck() error {
	return s.db.Ping()
}

func (s *DatabaseConnection) Shutdown() error {
	log.Print("database service shutdown")

	return s.entClient.Close()
}
