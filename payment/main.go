package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define a route for the payment service
	// This route will handle all requests starting with /payment/
	http.HandleFunc("/payment/", func(w http.ResponseWriter, r *http.Request) {
		// Respond with a message indicating the payment service is handling the request
		// The URL path is included in the response for debugging purposes
		fmt.Fprintf(w, "Payment Service: %s", r.URL.Path)
	})

	// Log a message indicating the payment service is running
	// This helps in identifying that the service has started successfully
	fmt.Println("Payment service running on port 8003")

	// Start the HTTP server on port 8003
	// This will block the main goroutine and listen for incoming requests
	http.ListenAndServe(":8003", nil)
}
