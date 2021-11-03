package service

import (
	"context"
	"errors"
	"github.com/ryan-holcombe/testable-golang/client"
	"github.com/ryan-holcombe/testable-golang/dao"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindAll(t *testing.T) {
	ctx := context.Background()
	var userID1 int64 = 1
	var userID2 int64 = 2
	users := []dao.UserModel{
		{
			ID:   userID1,
			Name: "Name 1",
		},
		{
			ID:   userID2,
			Name: "Name 2",
		},
	}
	tickets := []client.Ticket{
		{
			Barcode: "barcode",
		},
	}
	t.Run("error retrieving users", func(t *testing.T) {
		userDAO := &mockUserDAO{}
		inventoryClient := &mockInventoryClient{}
		unit := NewUserTicketsService(userDAO, inventoryClient)

		userDAO.On("FindAll", ctx).Return(nil, errors.New("timeout"))

		_, err := unit.FindAll(ctx)
		assert.Error(t, err)
	})

	t.Run("fail fast on ticket lookup failures", func(t *testing.T) {
		userDAO := &mockUserDAO{}
		inventoryClient := &mockInventoryClient{}
		unit := NewUserTicketsService(userDAO, inventoryClient)

		userDAO.On("FindAll", ctx).Return(users, nil)
		inventoryClient.On("FindUserTickets", userID1).Return(nil, errors.New("timeout"))

		_, err := unit.FindAll(ctx)
		assert.Error(t, err)
	})

	t.Run("successfully return users with tickets", func(t *testing.T) {
		userDAO := &mockUserDAO{}
		inventoryClient := &mockInventoryClient{}
		unit := NewUserTicketsService(userDAO, inventoryClient)

		userDAO.On("FindAll", ctx).Return(users, nil)
		inventoryClient.On("FindUserTickets", userID1).Return(tickets, nil)
		inventoryClient.On("FindUserTickets", userID2).Return(tickets, nil)

		results, err := unit.FindAll(ctx)
		assert.NoError(t, err)
		assert.Len(t, results, 2)
		assert.Equal(t, tickets, results[0].Tickets)
		assert.Equal(t, tickets, results[1].Tickets)
	})
}
