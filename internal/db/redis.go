package db

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type RedisStore struct {
	Client *redis.Client
}

func NewRedisStore(addr string) *RedisStore {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr, // "localhost:6379"
		Password: "",   // no password set
		DB:       0,    // use default DB
	})

	return &RedisStore{Client: rdb}
}

func (r *RedisStore) SaveOrderBook(symbol string, bids, asks [][]string) error {
	key := fmt.Sprintf("orderbook:%s", symbol)

	data := map[string]interface{}{
		"bids": bids,
		"asks": asks,
		"ts":   time.Now().UnixMilli(),
	}

	return r.Client.HSet(ctx, key, data).Err()
}
