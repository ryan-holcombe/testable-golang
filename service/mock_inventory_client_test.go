// Code generated by mockery. DO NOT EDIT.

package service

import (
	client "github.com/ryan-holcombe/testable-golang/client"
	mock "github.com/stretchr/testify/mock"
)

// mockInventoryClient is an autogenerated mock type for the inventoryClient type
type mockInventoryClient struct {
	mock.Mock
}

// FindUserTickets provides a mock function with given fields: userID
func (_m *mockInventoryClient) FindUserTickets(userID int64) ([]client.Ticket, error) {
	ret := _m.Called(userID)

	var r0 []client.Ticket
	if rf, ok := ret.Get(0).(func(int64) []client.Ticket); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]client.Ticket)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}