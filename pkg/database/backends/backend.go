package backends

import (
	"github.com/csyezheng/memcard/pkg/utils"
	"log"
)

// Backend is the database backend interface that must be implemented by all database backend.
type Backend interface {
	GetEngine() string
	DSN() string
}

func DefaultBackend() Backend {
	engine := utils.GetEnv("DB_ENGINE", "postgresql")
	var backend Backend
	switch engine {
	case "postgresql":
		backend = NewPostgresql()
	default:
		log.Fatalln("Unsupported database backend，only postgresql supported")
	}
	return backend
}
