//  Logger Middleware hocche ekta “Request Spy”
// Ja proti request er path, method, time, and context silently track kore — jeno tumi bujhte paro:
// ❓ Ke request pathacche
// ❓ Ki route e jachhe
// ❓ Kon method (GET, POST, etc.)
// ❓ Koto time lagchhe process korte
// ❓ Bypass hocche naki error hocche
// ❓ Suspicious attempt ache naki?

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Printf("🛰️ Incoming Request - Path: %s | Method: %s", r.URL.Path, r.Method)

		next.ServeHTTP(w, r)

		log.Printf("✅ Completed in %v\n", time.Since(start))
	})
}
