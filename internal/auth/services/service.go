package services

import (
	"fmt"
	authdb "github.com/csyezheng/memcard/internal/auth/database"
	"github.com/csyezheng/memcard/pkg/configs"
	"github.com/csyezheng/memcard/pkg/database"
	"github.com/csyezheng/memcard/pkg/httputils"
	"github.com/csyezheng/memcard/pkg/serverenv"
	"net/http"
)

type Service struct {
	config *configs.AuthServiceConfig
	env    *serverenv.ServerEnv
	db     *database.Database
	authDB *authdb.AuthDB
	JSON   func(http.ResponseWriter, int, interface{})
}

// NewService create a new authentication service
func NewService(config *configs.AuthServiceConfig, env *serverenv.ServerEnv) (*Service, error) {
	if env.Database() == nil {
		return nil, fmt.Errorf("missing database in server environment")
	}
	db := env.Database()
	authDB := authdb.NewAuthDB(db)
	return &Service{
		config: config,
		env:    env,
		db:     db,
		authDB: authDB,
		JSON:   httputils.JsonResponse,
	}, nil
}
