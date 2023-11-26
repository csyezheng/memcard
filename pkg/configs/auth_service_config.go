package configs

import (
	"fmt"
	"github.com/caarlos0/env/v10"
	"github.com/csyezheng/memcard/pkg/database/backends"
)

type AuthServiceConfig struct {
	DatabaseBackend backends.Backend
	Debug           bool   `env:"AUTH_SERVICE_DEBUG"`
	Port            string `env:"AUTH_SERVICE_PORT"`
}

func NewAuthServiceConfig() *AuthServiceConfig {
	var config AuthServiceConfig
	if err := env.Parse(&config); err != nil {
		fmt.Printf("%+v\n", err)
	}
	config.DatabaseBackend = backends.DefaultBackend()
	return &config
}
