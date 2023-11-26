package database

import (
	"errors"
)

var (
	// ErrValidationFailed is the error returned when validation failed.
	// This is user defined error.
	ErrValidationFailed = errors.New("validation failed")
)
