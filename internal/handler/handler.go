package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/LaLoca1/binance-l2-collector/internal/db"
	"github.com/LaLoca1/binance-l2-collector/internal/parser"
)

func HandleDepthMessage(msg *parser.DepthUpdateMessage, redisStore *db.RedisStore) {
	data, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Error marshalling depth update: %v", err)
		return
	}

	key := fmt.Sprintf("depth:%s", msg.Symbol)
	err = redisStore.Client.Set(context.Background(), key, data, 0).Err()
	if err != nil {
		log.Printf("Error storing depth update to Redis: %v", err)
	}
}
