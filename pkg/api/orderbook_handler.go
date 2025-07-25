package api

import (
	"context" // helps manage things like timeouts & deadlines (needed for Redis Calls)
	"encoding/json" // Helps convert data from Go structs to JSON
	"net/http"
	"strings"
)

// Handler for /orderbook/{symbol} route. Can access anything inside API like Redis using api variable 
func (api *API) orderbookHandler(w http.ResponseWriter, r *http.Request) {
	// r.URL.path is something like /orderbook/BTCUSDT & remove /orderbook/ part using TrimPrefix 
	symbol := strings.TrimPrefix(r.URL.Path, "/orderbook/") 
	if symbol == "" {
		http.Error(w, "symbol not provided", http.StatusBadRequest)
		return
	}

	// the redis key is something like depth:BTCUSDT. This is how redis keys are standardized for order book data. 
	key := "depth:" + strings.ToUpper(symbol)

	// Here we are asking Redis if there is anything stored under the key? 
	// context.background() - a base context that says "no timeout, no cancel" 
	// .Get(...).Result() - performs the read and gives you 2 things:
	// 	- val: the data from Redis (JSON String) & err: if something went wrong 
	val, err := api.Redis.Client.Get(context.Background(), key).Result()
	if err != nil {
		http.Error(w, "Not found in Redis", http.StatusNotFound)
		return
	}

	// Decode the JSON from Redis 
	
	// declaring a variable of type map...
	// This is a map (like a dictionary or JSON object). Keys are strings (string), like "bids" or "asks" 
	// Values are inferface{} - meaning anything. Could be number, string, array etc 
	// Why we do this, Since Redis stores the order book data as a JSON string, want to decode it into a flexible data structure in GO. 
	// You don't know the exact shape, so you can decode into a generic map that can hold any type of JSON structure. 
	var parsed map[string]interface{}

	// json.Unmarshal(...) -> This functions converts JSON into Go variable. 
	// []byte(val) -> val is a string (fetched from Redis). json.Unmarshal expects a byte slice([]byte) to convert string 
	// &parsed -> you are passing a pointer to parsed (&parsed) so that Unmarshal can fill it with the decoded data. This is how Go functions update values passed in 
	if err := json.Unmarshal([]byte(val), &parsed); err != nil {
		http.Error(w, "Failed to decode Redis data", http.StatusInternalServerError)
		return
	}

	// Returns the JSON response
	// Sets the HTTP header to let the client know it's JSON. Send parsed order book (as JSON) back to the user. 
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(parsed)
}
