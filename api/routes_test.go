package api

import (
	"errors"
	"fmt"
	"github.com/ryan-holcombe/testable-golang/client"
	"github.com/ryan-holcombe/testable-golang/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleUserTickets(t *testing.T) {
	t.Run("service returns error", func(t *testing.T) {
		mux := http.NewServeMux()
		userTicketsServiceMock := &mockUserTicketsService{}
		userTicketsServiceMock.On("FindAll", mock.Anything).Return(nil, errors.New("timeout"))

		RegisterRoutes(mux, userTicketsServiceMock)
		ts := httptest.NewServer(mux)
		defer ts.Close()

		req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/userTickets", ts.URL), nil)

		resp, err := ts.Client().Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("ok", func(t *testing.T) {
		mux := http.NewServeMux()

		userTickets := []service.UserWithTickets{
			{
				Name: "user",
				Tickets: []client.Ticket{
					{
						Barcode: "barcode",
					},
				},
			},
		}
		userTicketsServiceMock := &mockUserTicketsService{}
		userTicketsServiceMock.On("FindAll", mock.Anything).Return(userTickets, nil)

		RegisterRoutes(mux, userTicketsServiceMock)
		ts := httptest.NewServer(mux)
		defer ts.Close()

		req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/userTickets", ts.URL), nil)

		resp, err := ts.Client().Do(req)
		userTicketsJSON, err := io.ReadAll(resp.Body)
		assert.NoError(t, resp.Body.Close())

		expectedBody := `[{"name":"user","tickets":[{"barcode":"barcode"}]}]`

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.JSONEq(t, expectedBody, string(userTicketsJSON))
	})
}
