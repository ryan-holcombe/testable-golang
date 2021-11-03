package client

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFindUserTickets(t *testing.T) {
	t.Run("returns internal server error", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer ts.Close()

		var userID int64 = 1

		unit, err := NewInventoryClient(ts.URL)
		assert.NoError(t, err)

		_, err = unit.FindUserTickets(userID)
		assert.Error(t, err)
	})

	t.Run("tickets returned and successfully parsed", func(t *testing.T) {
		var userID int64 = 1

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/api/users/1/tickets", r.URL.Path)
			userTicketsJSON := `[{"barcode": "barcode1"}, {"barcode": "barcode2"}]'`
			w.Write([]byte(userTicketsJSON))
		}))
		defer ts.Close()

		unit, err := NewInventoryClient(ts.URL)
		assert.NoError(t, err)

		tickets, err := unit.FindUserTickets(userID)
		assert.NoError(t, err)
		assert.Len(t, tickets, 2)
		assert.Equal(t, "barcode1", tickets[0].Barcode)
		assert.Equal(t, "barcode2", tickets[1].Barcode)

	})
}
