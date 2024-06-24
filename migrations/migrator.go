package migrations

import (
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose"
)

func Migrate(db *sqlx.DB, path string) error {
	goose.SetDialect("postgres")
	return goose.Up(db.DB, path)
}
