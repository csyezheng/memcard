package database

import (
	"database/sql"
	"errors"
)

// IsNotFound determines if an error is a record not found.
func IsNotFound(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}

// IsValidationError returns true if the error is a validation error else false
func IsValidationError(err error) bool {
	return errors.Is(err, ErrValidationFailed)
}
