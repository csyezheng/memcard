package setup

import (
	"context"
	"fmt"
	"github.com/csyezheng/memcard/pkg/database"
	"github.com/csyezheng/memcard/pkg/database/backends"
	"github.com/csyezheng/memcard/pkg/logging"
	"github.com/csyezheng/memcard/pkg/serverenv"
)

// Setup runs common initialization code for all servers.
func Setup(ctx context.Context, config interface{}) (*serverenv.ServerEnv, error) {

	logger := logging.FromContext(ctx)

	// Build a list of options to pass to the server env.
	var envHelpers []serverenv.EnvHelper

	// Set up the database connection.
	if backend, ok := config.(backends.Backend); ok {
		logger.Info("configuring database")
		db := database.NewDatabase(backend)
		err := db.Open()
		if err != nil {
			return nil, fmt.Errorf("unable to connect to database: %w", err)
		}
		// Add serverEnv setup function
		envHelpers = append(envHelpers, serverenv.WithDatabase(db))
	}

	return serverenv.New(ctx, envHelpers...), nil
}
