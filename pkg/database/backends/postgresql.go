package backends

import (
	//"context"
	"fmt"
	"github.com/caarlos0/env/v10"
)

type Postgresql struct {
	// the database engine, always should be postgresql
	Engine string `env:"DB_ENGINE"`
	Host   string `env:"DB_HOST"`
	Port   int    `env:"DB_PORT"`
	// the name of the database to use
	Name     string `env:"DB_NAME"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	SSLMode  string `env:"DB_SSLMODE"`
}

func NewPostgresql() *Postgresql {
	var postgresql Postgresql
	if err := env.Parse(&postgresql); err != nil {
		fmt.Printf("%+v\n", err)
	}
	return &postgresql
}

func (db Postgresql) GetEngine() string {
	return db.Engine
}

func (db Postgresql) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", db.User, db.Password, db.Host, db.Port, db.Name, db.SSLMode)
}
