package services

import "github.com/csyezheng/memcard/pkg/database"

type Service struct {
	db           *database.Database
}

// NewService create a new authentication service
func NewService(db *database.Database) *Service {
	return &Service{
		db: db,
	}
}