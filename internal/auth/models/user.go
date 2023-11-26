package models

import (
	"time"
)

type User struct {
	ID           uint      `json:"id"`
	Username     string    `json:"username"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Password     string    `json:"password"`
	Email        string    `json:"email"`
	IsActive     bool      `json:"is_active"`
	IsStaff      bool      `json:"is_staff"`
	IsSupperUser bool      `json:"is_supper_user"`
	DateJoined   time.Time `json:"date_joined"`
	LastLogin    time.Time `json:"last_login"`
}
