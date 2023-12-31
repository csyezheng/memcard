package main

import (
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"github.com/csyezheng/memcard/pkg/configs"
	"github.com/csyezheng/memcard/pkg/logging"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5"
	"log"
	"os/signal"
	"syscall"
	"time"
)

var (
	pathFlag         = flag.String("path", "migrations/", "path to migrations folder")
	migrationTimeout = flag.Duration("timeout", 15*time.Minute, "duration for migration timeout")
)

func main() {
	flag.Parse()

	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer done()

	logger := logging.NewLoggerFromEnv()
	ctx = logging.AttachLogger(ctx, logger)

	config := configs.DefaultConfig()
	engine := config.DatabaseBackend.GetEngine()
	dsn := config.DatabaseBackend.DSN()
	logger.Debug(engine)
	logger.Debug(dsn)

	var driverName string
	switch engine {
	case "postgresql":
		driverName = "postgres"
	}

	db, err := sql.Open(driverName, dsn)
	if err != nil {
		logger.Error(err.Error())
	}
	var driver database.Driver

	switch engine {
	case "postgresql":
		driver, err = postgres.WithInstance(db, &postgres.Config{})
	}

	if err != nil {
		log.Fatalln(err.Error())
	}

	dir := fmt.Sprintf("file://%s", *pathFlag)
	m, err := migrate.NewWithDatabaseInstance(dir, "postgres", driver)
	if err != nil {
		log.Fatalln("failed create migrate: %w", err)
	}

	m.LockTimeout = *migrationTimeout

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalln("failed run migrate: %w", err)
	}
	srcErr, dbErr := m.Close()
	if srcErr != nil {
		log.Fatalln("migrate source error: %w", srcErr)
	}
	if dbErr != nil {
		log.Fatalln("migrate database error: %w", dbErr)
	}
	logger.Info("finished running migrations")
}
