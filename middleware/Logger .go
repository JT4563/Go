package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// 🕵️ Logger Middleware — Logs every request's path, method, time taken
func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now() // ⏱️ Start time

		log.Printf("🛰️ Incoming Request - Path: %s | Method: %s", r.URL.Path, r.Method)

		next.ServeHTTP(w, r) // ✅ Actual handler call

		log.Printf("✅ Completed in %v\n", time.Since(start)) // ⌛ Time taken
	})
}

// 🌍 Simple Handler — Hello Route
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "🌍 Hello from Logger Middleware Server!")
}

// 🔗 Setup middleware chaining manually
func main() {
	// New ServeMux
	mux := http.NewServeMux()

	// Wrap helloHandler with logger middleware
	mux.Handle("/hello", logger(http.HandlerFunc(helloHandler)))

	// Start server
	fmt.Println("🚀 Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
