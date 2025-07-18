package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/LaLoca1/binance-l2-collector/internal/db"
	"github.com/LaLoca1/binance-l2-collector/internal/ws"
)

// Storing the WebSocket URL as constant string. This the Binance feed for BTC/USDT order
// book depth, updated every 100ms.
const streamURL = "wss://stream.binance.com:9443/ws/btcusdt@depth@100ms"

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// Setup Redis connection
	redisStore := db.NewRedisStore("redis:6379", "", 0)

	// Setup WebSocket client
	client := ws.NewClient(streamURL)

	if err := client.Connect(); err != nil {
		log.Fatalf("WebSocket connection failed: %v", err)
	}

	go func() {
		http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("ok"))
		})
		log.Println("Healthcheck endpoint running on :8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatalf("Healthcheck server failed: %v", err)
		}
	}()

	client.Listen(interrupt, redisStore)
}
