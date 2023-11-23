package backends

import (
	"fmt"
	"github.com/caarlos0/env/v10"
)

type Mysql struct {
	// the database engine, always should be mysql
	Engine string `env:"DB_ENGINE"`
	Host   string `env:"DB_HOST"`
	Port   int    `env:"DB_PORT"`
	// the name of the database to use
	Name     string `env:"DB_NAME"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	SSLMode  string `env:"DB_SSLMODE"`
}

func (db Mysql) GetEngine() string {
	return db.Engine
}

func (db Mysql) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&tls=%s",
		db.User, db.Password, db.Host, db.Port, db.Name, db.SSLMode)
}

func NewMysql() *Mysql {
	mysql := Mysql{}
	if err := env.Parse(&mysql); err != nil {
		fmt.Printf("%+v\n", err)
	}
	return &mysql
}
