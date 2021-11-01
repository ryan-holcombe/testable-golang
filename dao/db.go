package dao

import (
	"github.com/jmoiron/sqlx"
)

func NewPostgresDB(dataSourceName string) (*sqlx.DB, error) {
	return sqlx.Open("postgres", dataSourceName)
}
