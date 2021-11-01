package api

import "github.com/ryan-holcombe/testable-golang/service"

type userTicketsService interface {
	FindAll() ([]service.UserWithTickets, error)
}
