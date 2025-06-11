package main

import (
	"fmt"
	"net/http"
)

// ✅ CORS Middleware
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Allow access from any frontend origin (you can specify exact domain later)
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Allow specific HTTP methods
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// Allow specific headers (needed for JWT, content-type etc.)
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight OPTIONS request
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Proceed to next middleware/handler
		next.ServeHTTP(w, r)
	})
}

// ✅ Sample Handler
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from backend!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)

	
	handlerWithCORS := corsMiddleware(mux)

	fmt.Println("🚀 Server running on http://localhost:8000")
	http.ListenAndServe(":8000", handlerWithCORS)
}
