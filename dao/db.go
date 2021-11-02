package dao

import (
	"github.com/jmoiron/sqlx"
)

// NewPostgresDB opens a new postgres connection
func NewPostgresDB(dataSourceName string) *sqlx.DB {
	return sqlx.MustOpen("postgres", dataSourceName)
}
