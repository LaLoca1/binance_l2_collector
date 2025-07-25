package api

import (
	"github.com/LaLoca1/binance-l2-collector/internal/db"
)

// API holds the dependencies needed by the HTTP server
type API struct {
	Redis *db.RedisStore
}

// NewAPI creates a new API instance with the Redis dependency injected.
func NewAPI(redis *db.RedisStore) *API {
	return &API{Redis: redis}
}
