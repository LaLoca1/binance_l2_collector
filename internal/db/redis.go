package db

import (
	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	Client *redis.Client
}

// NewRedisStore initializes a Redis connection
func NewRedisStore(addr, password string, db int) *RedisStore {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return &RedisStore{Client: rdb}
}
