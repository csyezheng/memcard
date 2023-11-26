package database

import (
	"github.com/csyezheng/memcard/pkg/database"
)

// AuthDB wraps a database connection and provides functions for authentication service.
type AuthDB struct {
	db *database.Database
}

// NewAuthDB creates a new AuthDB.
func NewAuthDB(db *database.Database) *AuthDB {
	return &AuthDB{
		db: db,
	}
}
