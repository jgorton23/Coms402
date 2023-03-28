package repo

import (
	"database/sql"

	sqladapter "github.com/Blank-Xu/sql-adapter"
	"github.com/casbin/casbin/v2/persist"
	"github.com/samber/do"

	_ "github.com/mattn/go-sqlite3"
)

func NewAuthorizationPolicyInMemoryRepo(i *do.Injector) (persist.Adapter, error) {
	// connect to the database first.
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	// Initialize an adapter and use it in a Casbin enforcer:
	// The adapter will use the Sqlite3 table name "casbin_rule_test",
	// the default table name is "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.
	a, err := sqladapter.NewAdapter(db, "sqlite3", "casbin_rule_test")
	if err != nil {
		return nil, err
	}

	return a, nil
}
