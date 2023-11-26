package database

import (
	"context"
	"database/sql"
	"github.com/csyezheng/memcard/pkg/database/backends"
	"github.com/csyezheng/memcard/pkg/logging"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"sync"
)

// Database is a handle to the database layer
type Database struct {
	db      *sql.DB
	dbLock  sync.Mutex
	backend backends.Backend
}

func NewDatabase(backend backends.Backend) *Database {
	return &Database{
		backend: backend,
	}
}

// Open creates a database connection. This should only be called once.
func (db *Database) Open() error {
	var rawDB *sql.DB
	var err error
	switch db.backend.GetEngine() {
	case "postgresql":
		dsn := db.backend.DSN()
		rawDB, err = sql.Open("pgx", dsn)
	default:
		log.Fatalf("not supported database engine: %s", db.backend.GetEngine())
	}
	if err != nil {
		log.Fatalf("failed to connect database:%s", err.Error())
	}
	db.db = rawDB
	return nil
}

// Close will close the database connection. Should be deferred right after Open.
func (db *Database) Close(ctx context.Context) error {
	logger := logging.FromContext(ctx)
	logger.Info("Closing database connection.")
	return db.db.Close()
}

func (db *Database) Execute(f func(db *sql.DB) error) error {
	return f(db.db)
}
