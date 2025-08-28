package parser

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/LaLoca1/binance-l2-collector/internal/db" // for working with Redis
)

// This function takes a parsed binance and update and stores it in redis
// msg is a pointer to a parsed depth update message from Binance.
// redisStore is a reference to the Redis wrapper that lets you store things in Redis.
func HandleDepthMessage(msg *DepthUpdateMessage, redisStore *db.RedisStore) {
	// Converts the msg (a Go struct) into a JSON byte array
	data, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Error marshalling depth update: %v", err)
		return
	}

	// Creates the Redis key where the message will be stored
	key := fmt.Sprintf("depth:%s", msg.Symbol)
	// Saves the JSON data to Redis using the Set command.
	// Context.Background() = no timeout, no cancellation. key = like "depth:BTCUSDT".
	// Data = JSON-formatted depth update, 0 = no expiration, .Err() = gets result of Redis command and checks if it failed
	err = redisStore.Client.Set(context.Background(), key, data, 0).Err()
	if err != nil {
		log.Printf("Error storing depth update to Redis: %v", err)
	}
}

// So this function
// 1) Takes a depth update
// 2) Turns it into JSON
// 3) Builds a Redis key like depth:BTCUSDT
// 4) Stores it in Redis
// 5) Logs any issues

// Example in action
// Lets say Binance sends a message for BTCUSDT, with bid/ask updates.
// The app will:
// - Parse it with parser.ParseDepthUpdate
// - Calls HandleDepthMessage(parsedMsg, redisStore)
// - The depth data is saved in Redis under depth:BTCUSDT
