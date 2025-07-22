// Declares that the file belongs to the API package. Then other parts of the app can import this package.
package api

// When dealing with HTTP requests & responses like creating a server, this is the go to package.
import (
	"net/http"
)

// HTTP Handler method for the /health route
// w http.ResponseWriter: Lets you write a http response back to the user 
// r *http.Request: Contains all details about the incoming HTTP request (method, URL, headers, etc.)
func (api *API) healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK) // Sets the HTTP status to 200 ok
	w.Write([]byte("ok"))        // Writes the response body - content shown when someone opens /health
}

// func (api *API) is a method receiver -> This will attach this function to the API struct, so it can use and access fields & methods on it
// api -> Name of receiver variable
// *API -> Type of receiver. Its a pointer to the API struct, meaning it has access to API fields (like Redis) & can modify them
// To summarise, a method called healthHandler is created * operates on the instance of *API
// Why do we use *API and not API? This is so that we don't copy the struct each time. Can access & mutate real underlying struct.

// If a value receiver is used like (api API):
// - The method operates on the original struct & any achanges to api.Redis inside the method won't affect the original. 

// Use pointer receivers when: 
// - The struct has fields that are pointers (like Redis clients)
// - Want to avoid copying big structs 
// - Might mutate the struct's state 