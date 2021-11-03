package dao

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestFindAll(t *testing.T) {
	t.Run("query error", func(t *testing.T) {
		ctx := context.Background()
		db := &mockSqlxDB{}
		unit := NewUserDAO(db)

		expectedQuery := "SELECT id, name FROM users"

		var results []UserModel
		db.On("SelectContext", ctx, &results, expectedQuery).Return(errors.New("invalid sql"))

		_, err := unit.FindAll(ctx)
		assert.Error(t, err)
	})

	t.Run("returns users", func(t *testing.T) {
		ctx := context.Background()
		db := &mockSqlxDB{}
		unit := NewUserDAO(db)

		expectedName1 := "name 1"
		expectedName2 := "name 2"
		expectedQuery := "SELECT id, name FROM users"

		// mock the DB select and append 2 results to the slice
		db.On("SelectContext", ctx, mock.AnythingOfType("*[]dao.UserModel"), expectedQuery).Return(nil).Run(func(args mock.Arguments) {
			dest := args.Get(1).(*[]UserModel)
			*dest = append(*dest, UserModel{
				ID:   1,
				Name: expectedName1,
			})
			*dest = append(*dest, UserModel{
				ID:   2,
				Name: expectedName2,
			})
		})

		results, err := unit.FindAll(ctx)
		assert.NoError(t, err)
		assert.Len(t, results, 2)
		assert.Equal(t, expectedName1, results[0].Name)
		assert.Equal(t, expectedName2, results[1].Name)
	})
}
