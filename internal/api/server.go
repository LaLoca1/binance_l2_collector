package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/LaLoca1/binance-l2-collector/internal/db"
)

// API holds the dependencies needed by the HTTP server
type API struct {
	Redis *db.RedisStore
}

// NewAPI creates a new API instance with the Redis dependency injected
func NewAPI(redis *db.RedisStore) *API {
	return &API{Redis: redis}
}

// StartServer sets up and runs the HTTP server
func (api *API) StartServer(addr string) {
	http.HandleFunc("/health", api.healthHandler)
	http.HandleFunc("/orderbook/", api.orderbookHandler)

	log.Printf("HTTP server running on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

// /health route
func (api *API) healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// /orderbook/{symbol}
func (api *API) orderbookHandler(w http.ResponseWriter, r *http.Request) {
	symbol := strings.TrimPrefix(r.URL.Path, "/orderbook/")
	if symbol == "" {
		http.Error(w, "symbol not provided", http.StatusBadRequest)
		return
	}

	key := "depth:" + strings.ToUpper(symbol)

	// Get raw JSON data from Redis
	val, err := api.Redis.Client.Get(context.Background(), key).Result()
	if err != nil {
		http.Error(w, "Not found in Redis", http.StatusNotFound)
		return
	}

	// Pretty print the response
	var parsed map[string]interface{}
	if err := json.Unmarshal([]byte(val), &parsed); err != nil {
		http.Error(w, "Failed to decode Redis data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(parsed)
}
