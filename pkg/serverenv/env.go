package serverenv

import (
	"context"
	"github.com/csyezheng/memcard/pkg/database"
)

// ServerEnv represents latent environment configuration for servers in this application.
type ServerEnv struct {
	database *database.Database
}

// EnvHelper defines function types to modify the ServerEnv on creation.
type EnvHelper func(*ServerEnv) *ServerEnv

// New creates a new ServerEnv with the requested options.
func New(ctx context.Context, helpers ...EnvHelper) *ServerEnv {
	env := &ServerEnv{}
	for _, f := range helpers {
		env = f(env)
	}
	return env
}

// WithDatabase attached a database to the environment.
func WithDatabase(db *database.Database) EnvHelper {
	return func(s *ServerEnv) *ServerEnv {
		s.database = db
		return s
	}
}

func (s *ServerEnv) Database() *database.Database {
	return s.database
}

// Close shuts down the server env, closing database connections, etc.
func (s *ServerEnv) Close(ctx context.Context) error {
	if s == nil {
		return nil
	}

	if s.database != nil {
		return s.database.Close(ctx)
	}

	return nil
}
