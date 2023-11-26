package main

import (
	"context"
	"github.com/csyezheng/memcard/internal/auth/routers"
	"github.com/csyezheng/memcard/internal/auth/services"
	"github.com/csyezheng/memcard/pkg/configs"
	"github.com/csyezheng/memcard/pkg/logging"
	"github.com/csyezheng/memcard/pkg/server"
	"github.com/csyezheng/memcard/pkg/setup"
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
	env, err := setup.Setup(ctx, backend)
	if err != nil {
		log.Fatalln("setup.Setup: %w", err)
	}
	defer env.Close(ctx)

	service, err := services.NewService(config, env)
	if err != nil {
		log.Fatalln(err)
	}

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
