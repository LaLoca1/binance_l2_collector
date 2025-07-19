package ws

import (
	"log"
	"os"
	"time"

	"github.com/LaLoca1/binance-l2-collector/internal/db"
	"github.com/LaLoca1/binance-l2-collector/internal/handler"
	"github.com/LaLoca1/binance-l2-collector/internal/parser"
	"github.com/gorilla/websocket"
)

// This defines a client struct
type Client struct {
	conn *websocket.Conn // A websocket connection object
	url  string          // Binance websocket URL this client connects to
}

// NewClient creates a new WebSocket client with the given URL
// A constructor function to create a new client. Takes in the WebSocket URL and returns a pointer
// to a new Client with url set and conn still nil.
func NewClient(url string) *Client {
	return &Client{url: url}
}

// Connect establishes the WebSocket connection to Binance
func (c *Client) Connect() error {
	// Tries to open a WebSocket connetion to the Binance stream URL stored in c.curl
	// Websocket.DefaultDialer.Dail initiates the connection. Returns 3 values: conn, _ httpresponse is not required, err:any connection error
	conn, _, err := websocket.DefaultDialer.Dial(c.url, nil)
	if err != nil {
		return err
	}
	// If successful, store the connection object in the Client struct and log a message.
	c.conn = conn
	log.Println("Connected to Binance WebSocket.")
	return nil
}

// Listen starts reading from the WebSocket and parses the depth messages
// Starts listening to incoming messages. Interrupt is a channel used to gracefully shut down (pass it from main.go)
func (c *Client) Listen(interrupt chan os.Signal, redisStore *db.RedisStore) {
	// done is a channel used to signal when the listener go routine ends
	done := make(chan struct{})
	// Starts a new goroutine (a lightweight thread) to read messages continuously in background.
	go func() {
		// Ensures that once this goroutine exits, the done channel is closed so main function knows its finished.
		defer close(done)
		for {
			// continously read messages from the websocket. First return value (message type) is ignored (_) because we only care
			// about message content. Message is the raw JSON payload sent by Binance.
			_, message, err := c.conn.ReadMessage()
			if err != nil {
				log.Printf("Read error: %v", err)
				return
			}

			// Call the custom parser function to convert the raw JSON into a structured DepthUpdateMessage GO struct.
			parsed, err := parser.ParseDepthUpdate(message)
			if err != nil {
				log.Printf("Parse error: %v", err)
				continue
			}
			// If parsing works, we send the parsed message to the handler which stores it in Redis. 
			handler.HandleDepthMessage(parsed, redisStore)
		}
	}()

	// Graceful shutdown logic
	// A blocking loop that waits for either:
	// - The listener goroutine to finish (via done) or the user to press CTRL+C (via interrupt)
	for {
		select {
		case <-done:
			return
		case <-interrupt:
			log.Println("Interrupt received, closing connection...")
			err := c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Printf("Close error: %v", err)
			}
			time.Sleep(time.Second)
			return
		}
	}
}

// Summary
// 1) Client Struct holds the connection and stream URL 
// 2) NewClient sets up a client (but doesn't connect yet)
// 3) Connect opens a WebSocket to Binance 
// 4) Listen does the heavy lifting:
// - Reads messages continously in a goroutine 
// - Parses them from JSON
// - Sends parsed messages to the handler