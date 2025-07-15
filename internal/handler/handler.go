package hander

import (
	"context"
	"log"

	"github.com/LaLoca1/binance-l2-collector/internal/parser"
	"github.com/redis/go-redis/v9"
)

func HandleDepthMessage(msg []byte, rdb *redis.Client) {
	depth, err := parser.ParseDepthUpdate(msg)
	if err != nil {
		log.Printf("Failed to parse message: %v", err)
		return
	}

	key := "orderbook:" + depth.Symbol + ":latest"
	err = rdb.HSet(context.Background(), key, map[string]interface{}{
		"bids": depth.Bids,
		"asks": depth.Asks,
	}).Err()
	if err != nil {
		log.Printf("Failed to store in Redis: %v", err)
		return
	}

	log.Println("Stored depth update in Redis for", depth.Symbol)
}
