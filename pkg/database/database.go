package database

import (
	"github.com/csyezheng/memcard/pkg/configs"
	"github.com/csyezheng/memcard/pkg/database/backends"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"sync"
)

// Database is a handle to the database layer
type Database struct {
	db      *gorm.DB
	dbLock  sync.Mutex
	backend backends.Backend
}

// Open creates a database connection. This should only be called once.
func (db *Database) Open() error {
	var rawDB *gorm.DB
	var err error
	switch db.backend.GetEngine() {
	case "postgresql":
		dsn := db.backend.DSN()
		rawDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	case "mysql":
		dsn := db.backend.DSN()
		rawDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
func (db *Database) Close() error {
	sqlDB, err := db.db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	return sqlDB.Close()
}

// LoadDatabase initialize database instance, it does not connect to the database.
func LoadDatabase() (*Database, error) {
	config := configs.DefaultConfig()
	backend := config.DatabaseBackend
	return &Database{
		backend: backend,
	}, nil
}
