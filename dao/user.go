package dao

import (
	"context"
)

// UserDAO access users in the DB
type UserDAO struct {
	db sqlxDB
}

// UserModel DB representation of the user
type UserModel struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

// NewUserDAO new instance of UserDAO
func NewUserDAO(db sqlxDB) *UserDAO {
	return &UserDAO{db}
}

// FindAll returns all the users
func (u UserDAO) FindAll(ctx context.Context) ([]UserModel, error) {
	var users []UserModel

	query := "SELECT id, name FROM users"
	if err := u.db.SelectContext(ctx, &users, query); err != nil {
		return nil, err
	}

	return users, nil
}
