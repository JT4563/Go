func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "🚫 Unauthorized: Missing Auth Header", http.StatusUnauthorized)
			return
		}
                           
//🧪 Example Request Header:

// Edit
// GET /secret HTTP/1.1
// Host: localhost:8080
// Authorization: Bearer abc.def.ghi
// Ekhane "Authorization" hocche header-er name, ar tar value hocche:

// Bearer abc.def.ghi

		                              
		// ✅ Optional: Bearer token check
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "🚫 Invalid Token Format", http.StatusUnauthorized)
			return
		}

		// 🪙 Token ke extract kora
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// TODO: JWT verification (advanced step)
		fmt.Println("✅ Token found:", token)

		next.ServeHTTP(w, r)
	})
}
