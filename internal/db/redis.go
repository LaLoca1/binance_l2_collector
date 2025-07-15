package storage

import (
	"github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // adjust if using Docker or Redis cloud
		Password: "",               // no password by default
		DB:       0,                // use default DB
	})
	return rdb
}
