package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new router using the Gorilla Mux package
	// The router will handle routing requests to the appropriate handlers
	r := mux.NewRouter()

	// Define a route for the /auth/ path prefix
	// Requests starting with /auth/ will be forwarded to the auth service running on port 8001
	r.PathPrefix("/auth/").Handler(reverseproxy("http://localhost:8001"))

	// Define a route for the /user/ path prefix
	// Requests starting with /user/ will be forwarded to the user service running on port 8002
	r.PathPrefix("/user/").Handler(reverseproxy("http://localhost:8002"))

	// Define a route for the /payment/ path prefix
	// Requests starting with /payment/ will be forwarded to the payment service running on port 8003
	r.PathPrefix("/payment/").Handler(reverseproxy("http://localhost:8003"))

	// Log a message indicating the API Gateway is running
	// This helps in identifying that the gateway has started successfully
	log.Println("API gateway running on 8080")

	// Start the HTTP server on port 8080 and use the router to handle requests
	// This will block the main goroutine and listen for incoming requests
	log.Fatal(http.ListenAndServe(":8080", r))
}

// reverseproxy creates a reverse proxy for a given target URL
// It forwards incoming requests to the target and sends back the response to the client
func reverseproxy(target string) http.Handler {
	// Parse the target URL into a URL object
	// This is required to create the reverse proxy
	url, _ := url.Parse(target)

	// Create a new single-host reverse proxy for the target URL
	proxy := httputil.NewSingleHostReverseProxy(url)

	// Return an HTTP handler that uses the reverse proxy to handle requests
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the path of the request being proxied
		// This helps in debugging and monitoring the requests being forwarded
		log.Println("🔁 Proxying:", r.URL.Path)

		// Use the reverse proxy to forward the request to the target service
		// The proxy will handle sending the request and returning the response
		proxy.ServeHTTP(w, r)
	})
}