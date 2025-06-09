// func logger(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		start := time.Now()

// 		log.Printf("🛰️ Incoming Request - Path: %s | Method: %s", r.URL.Path, r.Method)

// 		next.ServeHTTP(w, r)

// 		log.Printf("✅ Completed in %v\n", time.Since(start))
// 	})
// }
