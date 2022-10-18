package repo

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"
	"github.com/samber/do"

	"github.com/MatthewBehnke/exampleGoApi/internal/entity"
	"github.com/MatthewBehnke/exampleGoApi/pkg/database/ent"
)

func NewDatabaseService(i *do.Injector) (*DatabaseService, error) {
	dbService := &DatabaseService{}
	dbService.config = do.MustInvoke[*entity.Config](i)

	db, err := sql.Open("postgres", dbService.config.PG.URL)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to: %v", err)
	}

	dbService.db = db

	dbService.db.SetMaxOpenConns(dbService.config.PG.PoolMax)

	drv := entsql.OpenDB("postgres", dbService.db)

	dbService.entClient = ent.NewClient(ent.Driver(drv))

	// Run the auto migration tool.
	if err := dbService.entClient.Schema.Create(context.Background()); err != nil {
		return nil, fmt.Errorf("failed creating schema resources: %v", err)
	}
	log.Print("database service started")

	return dbService, nil
}

type DatabaseService struct {
	config    *entity.Config
	db        *sql.DB
	entClient *ent.Client
}

func (s *DatabaseService) Client() *ent.Client {
	return s.entClient
}

func (s *DatabaseService) HealthCheck() error {
	return s.db.Ping()
}

func (s *DatabaseService) Shutdown() error {
	log.Print("database service shutdown")
	return s.entClient.Close()
}
