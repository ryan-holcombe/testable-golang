package service

import (
	"context"
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
	Name    string          `json:"name"`
	Tickets []client.Ticket `json:"tickets"`
}

// FindAll returns all the users and their tickets
func (u UserTicketsService) FindAll(ctx context.Context) ([]UserWithTickets, error) {
	var results []UserWithTickets

	// retrieve all users
	users, err := u.userDAO.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	// iterate over all users and find their tickets
	for _, user := range users {
		tickets, err := u.inventoryClient.FindUserTickets(user.ID)
		if err != nil {
			return nil, err
		}
		results = append(results, UserWithTickets{
			Name:    user.Name,
			Tickets: tickets,
		})
	}
	return results, nil
}
