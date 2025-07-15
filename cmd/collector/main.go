package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/LaLoca1/binance-l2-collector/internal/ws"
)

// Storing the WebSocket URL as constant string. This the Binance feed for BTC/USDT order
// book depth, updated every 100ms. 
const streamURL = "wss://stream.binance.com:9443/ws/btcusdt@depth@100ms"

func main() {
	// Creating a channel that can hold os.Signal values. Will listen for interrupt signals. 
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// Calling constructor function from ws. Creates a new client and gives it websocket url. 
	client := ws.NewClient(streamURL)

	if err := client.Connect(); err != nil {
		log.Fatalf("WebSocket connection failed: %v", err)
	}

	// Now listening to WebSocket stream using custom client. 
	client.Listen(interrupt)
}
