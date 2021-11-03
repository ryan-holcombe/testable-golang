package api

import (
	"encoding/json"
	"net/http"
)

// RegisterRoutes registers API routes for this service
func RegisterRoutes(mux *http.ServeMux, userTicketsService userTicketsService) {
	mux.Handle("/api/userTickets", handleUserTickets(userTicketsService))
}

func handleUserTickets(service userTicketsService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		results, err := service.FindAll(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		jsonBytes, _ := json.Marshal(&results)
		_, _ = w.Write(jsonBytes)
	}
}
