package service

import (
	"github.com/ryan-holcombe/testable-golang/client"
)

// NewUserTicketsService creates a new UserTicketsService
func NewUserTicketsService(userDAO userDAO, inventoryClient inventoryClient) *UserTicketsService {
	return &UserTicketsService{
		userDAO:         userDAO,
		inventoryClient: inventoryClient,
	}
}

// UserTicketsService just an example handler
type UserTicketsService struct {
	userDAO         userDAO
	inventoryClient inventoryClient
}

// UserWithTickets model representing a user with a list of tickets
type UserWithTickets struct {
	Name    string
	Tickets []client.Ticket
}

// FindAll returns all the users and their tickets
func (t UserTicketsService) FindAll() ([]UserWithTickets, error) {
	return nil, nil
}
