package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Logger Middleware
func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Printf("Requested URL: %s | Method: %s", r.URL.Path, r.Method)

		next.ServeHTTP(w, r) // next handler ke call kora hocche

		log.Printf("Completed in %v", time.Since(start))
	})
}

// Real handler (final target)
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func main() {
	mux := http.NewServeMux()

	// /hello path e helloHandler ke attach kora hocche
	mux.Handle("/hello", logger(http.HandlerFunc(helloHandler)))

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
