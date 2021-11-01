package dao

import (
	"github.com/jmoiron/sqlx"
)

func NewPostgresDB(dataSourceName string) *sqlx.DB {
	return sqlx.MustOpen("postgres", dataSourceName)
}
