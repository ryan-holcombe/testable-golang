package dao

import (
	"context"
)

type sqlxDB interface {
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}
