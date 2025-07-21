package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/LaLoca1/binance-l2-collector/internal/api"
	"github.com/LaLoca1/binance-l2-collector/internal/db"
	"github.com/LaLoca1/binance-l2-collector/internal/ws"
)

// Storing the WebSocket URL as constant string. This the Binance feed for BTC/USDT order
// book depth, updated every 100ms.
const streamURL = "wss://stream.binance.com:9443/ws/btcusdt@depth@100ms"

func main() {
	// Creates a channel called interrupt that listens for OS signals
	interrupt := make(chan os.Signal, 1)
	// Tells Go to send os.Interrupt signals(like ctrl+c) into channel. This then
	// lets you gracefully shut down the WebSocket client later
	signal.Notify(interrupt, os.Interrupt)

	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379" // fallback
	}
	// This calls the custom function initialize a redis client
	// returns a *RedisStore struct with a redis client you can use throughout app
	redisStore := db.NewRedisStore(redisAddr, "", 0)

	// Calls the constructor NewClient(from internal ws package) with Binance URL
	// Sets sets the websocket url field in the client. its just created - not connected yet.
	client := ws.NewClient(streamURL)

	// Tries to open websocket connection to Binance. Client.Connect() dials the Binance server
	// and stores the connection in the client.
	if err := client.Connect(); err != nil {
		log.Fatalf("WebSocket connection failed: %v", err)
	}

	// Start HTTP API server (includes /health and /orderbook routes)
	apiServer := api.NewAPI(redisStore)
	go apiServer.StartServer(":8080")

	// Interrupt channel -> so it knows when the shut down
	// RedisStore -> so it can store parsed messages in Redis
	// This will read binance messages, parses them, stores in redis
	client.Listen(interrupt, redisStore)
}
