package db

import (
	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	Client *redis.Client // A pointer to the Redis Client object that knows how to talk to your Redis Server
}

// A constructor function. Creates and returns a new RedisStore.
func NewRedisStore(addr, password string, db int) *RedisStore {
	// This lines creates a new Redis client using the settings passed in.
	// This results in a usable Redis connection rdb
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,     // Where to connect
		Password: password, // How to authenticate
		DB:       db,       // Which Redis DB to use
	})
	// This builds and returns a new RedisStore struct, putting your redis client (rdb) inside it
	// The & means your returning a pointer to that struct, so:
	// - Avoid copying the whole struct & can use the same connection throughout the app.
	return &RedisStore{Client: rdb}
}

// RedisStore - App's wrapper for working with Redis
// NewRedisStore(..) - Connects to Redis and gives you usable client
