package repo

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/samber/do"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

// Pattern to verify DatabaseEnt conforms to the required interfaces
var (
	_assertDatabaseEntImplem                    = &DatabaseEnt{}
	_                        do.Shutdownable    = _assertDatabaseEntImplem
	_                        do.Healthcheckable = _assertDatabaseEntImplem
)

func DatabaseEntBuilder() *databaseEntBuilder {
	return &databaseEntBuilder{
		opts: struct{ maxOpenConns int }{
			maxOpenConns: 10,
		},
	}
}

type databaseEntBuilder struct {
	opts struct {
		maxOpenConns int
	}
}

func (builder *databaseEntBuilder) WithMaxOpenConns(maxConns int) *databaseEntBuilder {
	builder.opts.maxOpenConns = maxConns
	return builder
}

func (builder *databaseEntBuilder) WithPostgres(dsn string) (*DatabaseEnt, error) {
	return builder.dbSetup("postgres", dsn)
}

func (builder *databaseEntBuilder) WithSqlite(dsn string) (*DatabaseEnt, error) {
	return builder.dbSetup("sqlite", dsn)
}

func (builder *databaseEntBuilder) dbSetup(dbType string, dsn string) (*DatabaseEnt, error) {
	db, err := sql.Open(dbType, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to: %w", err)
	}

	db.SetMaxOpenConns(builder.opts.maxOpenConns)

	drv := entsql.OpenDB(dbType, db)

	dbImplem := &DatabaseEnt{
		client: ent.NewClient(ent.Driver(drv)),
		db:     db,
	}

	if err := dbImplem.client.Schema.Create(context.Background()); err != nil {
		return nil, fmt.Errorf("failed creating schema resources: %w", err)
	}

	log.Print("database service started")

	return dbImplem, nil
}

type DatabaseEnt struct {
	db     *sql.DB
	client *ent.Client
}

func (implem *DatabaseEnt) HealthCheck() error {
	return implem.db.Ping()
}

func (implem *DatabaseEnt) Shutdown() error {
	return implem.client.Close()
}
