package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// InventoryClient inventory HTTP client
type InventoryClient struct {
	baseURL *url.URL
}

// Ticket user ticket
type Ticket struct {
	Barcode string `json:"barcode"`
}

// NewInventoryClient creates a new InventoryClient from a base URL
func NewInventoryClient(baseURL string) (*InventoryClient, error) {
	parsedURL, err := url.Parse(baseURL)
	return &InventoryClient{
		baseURL: parsedURL,
	}, err
}

// FindUserTickets retrieve the tickets from a user ID
func (i InventoryClient) FindUserTickets(userID int64) ([]Ticket, error) {
	userTicketURL, _ := i.baseURL.Parse(fmt.Sprintf("/api/users/%d/tickets", userID))
	resp, err := http.DefaultClient.Get(userTicketURL.String())
	if err != nil {
		return nil, err
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid http status [%d] received", resp.StatusCode)
	}
	defer resp.Body.Close()

	var results []Ticket
	err = json.NewDecoder(resp.Body).Decode(&results)
	return results, err
}
