package api

import (
	"net/http"
)

// RegisterRoutes registers API routes for this service
func RegisterRoutes(mux *http.ServeMux, userTicketsService userTicketsService) {
	mux.Handle("/", handleTestEndpoint(userTicketsService))
}

func handleTestEndpoint(service userTicketsService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}
