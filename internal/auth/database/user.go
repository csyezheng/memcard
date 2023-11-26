package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/csyezheng/memcard/internal/auth/models"
	"github.com/csyezheng/memcard/pkg/database"
	"github.com/csyezheng/memcard/pkg/logging"
)

func (adb *AuthDB) FindUserByEmailAndUsername(ctx context.Context, username string, email string) (*models.User, error) {
	logger := logging.FromContext(ctx)
	query := `
		SELECT 
			username, email
		FROM 
			auth.users
		WHERE 
			username = $1 OR email = $2;
`
	user := models.User{}
	err := adb.db.Execute(func(db *sql.DB) error {
		row := db.QueryRowContext(ctx, query, username, email)
		if err := row.Scan(&user.Username, &user.Email); err != nil {
			if database.IsNotFound(err) {
				return err
			}
			logger.Error(err.Error())
			return fmt.Errorf("scanning results: %w", err)
		}
		return nil
	})

	return &user, err
}

func (adb *AuthDB) SaveUser(ctx context.Context, user *models.User) error {
	logger := logging.FromContext(ctx)
	query := `
		INSERT INTO
			auth.users
			(username, first_name, last_name, password, email, date_joined)
		VALUES
			($1, $2, $3, $4, $5, $6);
`
	err := adb.db.Execute(func(db *sql.DB) error {
		_, err := db.ExecContext(ctx, query, user.Username, user.FirstName, user.LastName, user.Password, user.Email, user.DateJoined)
		return err
	})

	if err != nil {
		logger.Error("failed to save user", "Error", err)
		return err
	}
	return nil
}
