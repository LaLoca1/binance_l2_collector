// This file defines how to parse JSON messages received from Binance's WebSocket into a Go struct,
// specifically depth updates(i.e changes to the order book)

// Declares that this file is part of the parser package. Makes it reusable by other parts of the app.
package parser

// used for parsing(unmarshalling) JSON into Go Structs
import "encoding/json"

// A struct is defined that matches the structure of Binance's depth update messages.
// Each field is tagged with json:"...", which tells Go how to map JSON fields to Go fields
type DepthUpdateMessage struct {
	EventType string     `json:"e"` // Type of event (depthupdate)
	EventTime int64      `json:"E"` // Time event was generated
	Symbol    string     `json:"s"` // Trading pair symbol (BTCUSDT)
	Bids      [][]string `json:"b"` // List of bid updates: [[price, quantity], ...]
	Asks      [][]string `json:"a"` // List of ask updates: [[price, quantity], ...]
}

// This defines a function that takes in a raw byte slice ([] byte) - usually result of ReadMessage() - and returns either:
// - A pointer to a DepthUpdateMessage struct if parsing succeeded
// - Or an error if JSON decoding failed
func ParseDepthUpdate(data []byte) (*DepthUpdateMessage, error) {

	// Declares a variable named 'msg' of type 'DepthUpdateMessage'. This is where the parsed JSON will be stored.
	var msg DepthUpdateMessage

	// The line attempts to unmarshal the 'data' JSON into the 'msg' struct:
	// - 'data' is the input raw JSON
	// - '&msg' is a pointer so Go can modify it.
	// - If unmarshalling fails (e.g invalid JSON), it returns the error.
	if err := json.Unmarshal(data, &msg); err != nil {
		return nil, err
	}

	// If successful, it returns:
	// - A pointer to the filled-in msg struct.
	// nil error to indicate success
	return &msg, nil
}
