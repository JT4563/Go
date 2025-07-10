package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define a route for the user service
	// This route will handle all requests starting with /user/
	http.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {
		// Respond with a message indicating the user service is handling the request
		// The URL path is included in the response for debugging purposes
		fmt.Fprintf(w, "User Service: %s", r.URL.Path)
	})

	// Log a message indicating the user service is running
	// This helps in identifying that the service has started successfully
	fmt.Println("User service running on port 8002")

	// Start the HTTP server on port 8002
	// This will block the main goroutine and listen for incoming requests
	http.ListenAndServe(":8002", nil)
}
