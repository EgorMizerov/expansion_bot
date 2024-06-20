package database

import (
	"fmt"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type ConnectConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func ConnectToDatabase(config ConnectConfig) (*sqlx.DB, error) {
	url := getConnectionUrl(config)
	conn, err := sqlx.Connect("pgx", url)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func getConnectionUrl(cfg ConnectConfig) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.Password, cfg.Password, cfg.Database)
}
