package service

import (
	"context"
	"github.com/ryan-holcombe/testable-golang/client"
	"github.com/ryan-holcombe/testable-golang/dao"
)

type userDAO interface {
	FindAll(ctx context.Context) ([]dao.UserModel, error)
}

type inventoryClient interface {
	FindUserTickets(userID int64) ([]client.Ticket, error)
}
