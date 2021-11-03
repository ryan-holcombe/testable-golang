package api

import (
	"context"
	"github.com/ryan-holcombe/testable-golang/service"
)

type userTicketsService interface {
	FindAll(ctx context.Context) ([]service.UserWithTickets, error)
}
