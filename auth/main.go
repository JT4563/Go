package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define a route for the auth service
	// This route will handle all requests starting with /auth/
	http.HandleFunc("/auth/", func(w http.ResponseWriter, r *http.Request) {
		// Respond with a message indicating the auth service is handling the request
		// The URL path is included in the response for debugging purposes
		fmt.Fprintf(w, "Auth Service: %s", r.URL.Path)
	})

	// Log a message indicating the auth service is running
	// This helps in identifying that the service has started successfully
	fmt.Println("Auth service running on port 8001")

	// Start the HTTP server on port 8001
	// This will block the main goroutine and listen for incoming requests
	http.ListenAndServe(":8001", nil)
}
