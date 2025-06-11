package main

import (
	"fmt"
	"net/http"

	"github.com/rs/cors"
)

// ✅ Sample route handler
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Go backend! 🌍")
}

func main() {
	// ✅ Create your ServeMux (router)
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)

	// ✅ Configure CORS options (production safe)
	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Frontend domain
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// ✅ Wrap your router with the CORS middleware
	handler := corsOptions.Handler(mux)

	// ✅ Start server
	fmt.Println("🚀 Server running on http://localhost:8000")
	http.ListenAndServe(":8000", handler)
}
