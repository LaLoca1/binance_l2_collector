package api

import (
	"log"
	"net/http"
)

// StartServer sets up and runs the HTTP Server
func (api *API) StartServer(addr string) {
	http.HandleFunc("/health", api.healthHandler)
	http.HandleFunc("/orderbook/", api.orderbookHandler)

	log.Printf("HTTP server running on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

// so for the healhHandler endpoint, you pass api.healthHandler(Method bound to API struct) as function to handle /health requests
// When a request hits /health, gets routed to your method, and api lets handler access shared dependencies like Redis. 