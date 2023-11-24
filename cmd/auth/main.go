package main

import (
	"context"
	"github.com/csyezheng/memcard/internal/auth/routers"
	"github.com/csyezheng/memcard/internal/auth/services"
	"github.com/csyezheng/memcard/pkg/configs"
	"github.com/csyezheng/memcard/pkg/database"
	"github.com/csyezheng/memcard/pkg/logging"
	"github.com/csyezheng/memcard/pkg/server"
	"log"
	"os/signal"
	"syscall"
)

func main() {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	logger := logging.NewLoggerFromEnv()
	ctx = logging.AttachLogger(ctx, logger)

	config := configs.NewAuthServiceConfig()

	// Setup database
	backend := config.DatabaseBackend
	db, err := database.LoadDatabase(backend)
	if err != nil {
		log.Fatalln("failed to load database config: %w", err)
	}
	if err := db.Open(); err != nil {
		log.Fatalln("failed to connect to database: %w", err)
	}
	defer db.Close()

	service := services.NewService(db)
	router := routers.RegisterRoutes(service)

	srv, err := server.NewServer(config.Port)
	if err != nil {
		log.Fatalln(err)
	}

	err = srv.ServeHTTPHandler(ctx, router)
	if err != nil {
		log.Fatalln(err)
	}

	logger.Info("Successful shutdown server")
}