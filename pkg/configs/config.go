package configs

import (
	"github.com/csyezheng/memcard/pkg/database/backends"
)

type Config struct {
	DatabaseBackend backends.Backend
}

func DefaultConfig() *Config {
	return &Config{
		DatabaseBackend: backends.DefaultBackend(),
	}
}
