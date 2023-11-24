package configs

import (
	"github.com/csyezheng/memcard/pkg/database/backends"
	"github.com/csyezheng/memcard/pkg/utils"
	"log"
)

type Config struct {
	DatabaseBackend backends.Backend
}

func DefaultConfig() *Config {
	engine := utils.GetEnv("DB_ENGINE", "postgresql")
	var backend backends.Backend
	switch engine {
	case "postgresql":
		backend = backends.NewPostgresql()
	case "mysql":
		backend = backends.NewMysql()
	default:
		log.Fatalln("Unsupported database backend，only mysql, postgresql supported")
	}
	return &Config{
		DatabaseBackend: backend,
	}
}
