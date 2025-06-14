// comprehensive_middleware.go
// A production-ready middleware implementation that combines all middleware concepts
package middleware

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"strings"
	"sync"
	"time"
)

// Middleware represents a function that wraps an http.Handler and returns a new http.Handler
type Middleware func(http.Handler) http.Handler

// ==================== CORE MIDDLEWARE SYSTEM ====================

// Chain combines multiple middleware into a single middleware
// Middleware will be executed in the order they're provided (first to last)
func Chain(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		// Apply middleware in reverse order so they execute in the order provided
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}
		return next
	}
}

// Apply attaches all provided middleware to an http.Handler
func Apply(h http.Handler, middlewares ...Middleware) http.Handler {
	for _, m := range middlewares {
		h = m(h)
	}
	return h
}

// StandardStack returns a chain of commonly used middleware
func StandardStack() Middleware {
	return Chain(
		Recovery(),
		Logger(),
		CORS(),
	)
}

// FullStack returns a chain of all available middleware
func FullStack() Middleware {
	return Chain(
		Recovery(),
		Logger(),
		CORS(),
		RateLimit(100),
		BasicAuth("admin", "password"), // For demonstration - use secure credentials in production
	)
}

// ==================== LOGGING MIDDLEWARE ====================

// responseWriter is a wrapper for http.ResponseWriter that captures status and bytes
type responseWriter struct {
	http.ResponseWriter
	status int
	bytes  int
}

func (rw *responseWriter) WriteHeader(status int) {
	rw.status = status
	rw.ResponseWriter.WriteHeader(status)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	n, err := rw.ResponseWriter.Write(b)
	rw.bytes += n
	return n, err
}

// LoggerConfig holds configuration for the Logger middleware
type LoggerConfig struct {
	// LogFunc is called for each request with timing and status info
	LogFunc func(req *http.Request, status, bytes int, duration time.Duration)
	// ExcludePaths lists URL paths that won't be logged
	ExcludePaths []string
}

// DefaultLoggerConfig returns a default logger configuration
func DefaultLoggerConfig() LoggerConfig {
	return LoggerConfig{
		LogFunc: func(req *http.Request, status, bytes int, duration time.Duration) {
			log.Printf(
				"%s %s %s - %d %d - %v",
				req.RemoteAddr,
				req.Method,
				req.URL.Path,
				status,
				bytes,
				duration,
			)
		},
		ExcludePaths: []string{"/health", "/metrics"},
	}
}

// Logger returns a middleware that logs HTTP requests
func Logger() Middleware {
	return LoggerWithConfig(DefaultLoggerConfig())
}

// LoggerWithConfig returns a middleware that logs HTTP requests with custom config
func LoggerWithConfig(config LoggerConfig) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Skip logging for excluded paths
			for _, path := range config.ExcludePaths {
				if r.URL.Path == path {
					next.ServeHTTP(w, r)
					return
				}
			}

			// Create a response writer that captures status and byte count
			rw := &responseWriter{
				ResponseWriter: w,
				status:         http.StatusOK,
			}

			start := time.Now()
			next.ServeHTTP(rw, r)
			duration := time.Since(start)

			config.LogFunc(r, rw.status, rw.bytes, duration)
		})
	}
}

// ==================== RECOVERY MIDDLEWARE ====================

// RecoveryConfig holds configuration for the Recovery middleware
type RecoveryConfig struct {
	// OnPanic is called when a panic occurs with the error and stack trace
	OnPanic func(err interface{}, stack []byte)
	// ResponseStatus is the HTTP status code to return after recovering from a panic
	ResponseStatus int
	// ResponseBody is the response body to return after recovering from a panic
	ResponseBody string
}

// DefaultRecoveryConfig returns a default recovery configuration
func DefaultRecoveryConfig() RecoveryConfig {
	return RecoveryConfig{
		OnPanic: func(err interface{}, stack []byte) {
			log.Printf("PANIC: %v\n%s", err, string(stack))
		},
		ResponseStatus: http.StatusInternalServerError,
		ResponseBody:   "Internal Server Error",
	}
}

// Recovery returns a middleware that recovers from panics
func Recovery() Middleware {
	return RecoveryWithConfig(DefaultRecoveryConfig())
}

// RecoveryWithConfig returns a middleware that recovers from panics with custom config
func RecoveryWithConfig(config RecoveryConfig) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					stack := debug.Stack()
					config.OnPanic(err, stack)

					w.WriteHeader(config.ResponseStatus)
					w.Write([]byte(config.ResponseBody))
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}

// ==================== CORS MIDDLEWARE ====================

// CORSConfig holds configuration for the CORS middleware
type CORSConfig struct {
	// AllowOrigins is a list of origins a cross-domain request can be executed from
	AllowOrigins []string
	// AllowMethods is a list of methods allowed to be used in CORS requests
	AllowMethods []string
	// AllowHeaders is a list of headers allowed in CORS requests
	AllowHeaders []string
	// AllowCredentials indicates whether the request can include user credentials
	AllowCredentials bool
	// MaxAge indicates how long the results of a preflight request can be cached
	MaxAge int
}

// DefaultCORSConfig returns a default CORS configuration
func DefaultCORSConfig() CORSConfig {
	return CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		MaxAge:           86400, // 24 hours
	}
}

// CORS returns a middleware that handles CORS requests
func CORS() Middleware {
	return CORSWithConfig(DefaultCORSConfig())
}

// CORSWithConfig returns a middleware that handles CORS requests with custom config
func CORSWithConfig(config CORSConfig) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Set CORS headers
			origin := r.Header.Get("Origin")
			
			// Check if origin is allowed
			allowOrigin := "*"
			if len(config.AllowOrigins) > 0 && config.AllowOrigins[0] != "*" {
				allowOrigin = ""
				for _, o := range config.AllowOrigins {
					if o == origin {
						allowOrigin = origin
						break
					}
				}
			}
			
			if allowOrigin != "" {
				w.Header().Set("Access-Control-Allow-Origin", allowOrigin)
			}

			// Set other CORS headers
			w.Header().Set("Access-Control-Allow-Methods", strings.Join(config.AllowMethods, ", "))
			w.Header().Set("Access-Control-Allow-Headers", strings.Join(config.AllowHeaders, ", "))
			
			if config.AllowCredentials {
				w.Header().Set("Access-Control-Allow-Credentials", "true")
			}
			
			if config.MaxAge > 0 {
				w.Header().Set("Access-Control-Max-Age", fmt.Sprintf("%d", config.MaxAge))
			}

			// Handle OPTIONS preflight requests
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// ==================== AUTHENTICATION MIDDLEWARE ====================

// BasicAuth returns a middleware that enforces HTTP Basic Authentication
func BasicAuth(username, password string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get credentials
			user, pass, ok := r.BasicAuth()

			if !ok || user != username || pass != password {
				w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// JWTConfig holds configuration for the JWT middleware
type JWTConfig struct {
	// Secret used to verify the token
	Secret string
	// TokenLookup is a string in the form of "<source>:<name>" that is used
	// to extract the token from the request. 
	// Use "header:Authorization", "query:token", "cookie:jwt"
	TokenLookup string
	// ErrorHandler is a function to handle JWT errors
	ErrorHandler func(w http.ResponseWriter, r *http.Request, err error)
}

// JWT returns a middleware that validates JWT tokens
// Note: This is a simplified implementation for demonstration purposes
func JWT(secret string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get token from Authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Authorization token required", http.StatusUnauthorized)
				return
			}

			// Strip 'Bearer ' prefix
			tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

			// In a real application, validate the token here using a JWT library
			// For this example, we'll do a simple check if the token matches the secret
			if tokenString != secret {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// ==================== RATE LIMITING MIDDLEWARE ====================

// RateLimiter implements a simple in-memory rate limiter
type RateLimiter struct {
	requests map[string][]time.Time
	limit    int
	duration time.Duration
	mu       sync.Mutex
}

// NewRateLimiter creates a new rate limiter instance
func NewRateLimiter(limit int, duration time.Duration) *RateLimiter {
	return &RateLimiter{
		requests: make(map[string][]time.Time),
		limit:    limit,
		duration: duration,
	}
}

// Allow checks if a request is allowed based on the rate limit
func (r *RateLimiter) Allow(key string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	
	// Remove old timestamps
	var recent []time.Time
	for _, timestamp := range r.requests[key] {
		if now.Sub(timestamp) <= r.duration {
			recent = append(recent, timestamp)
		}
	}
	
	r.requests[key] = recent
	
	// Check if under limit
	if len(recent) < r.limit {
		r.requests[key] = append(r.requests[key], now)
		return true
	}
	
	return false
}

// RateLimit returns a middleware that limits requests per client
func RateLimit(requestsPerMinute int) Middleware {
	limiter := NewRateLimiter(requestsPerMinute, time.Minute)
	
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Use client IP as the rate limit key
			key := r.RemoteAddr
			
			if !limiter.Allow(key) {
				http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
				return
			}
			
			next.ServeHTTP(w, r)
		})
	}
}

// ==================== EXAMPLE USAGE ====================

/*
func ExampleUsage() {
	// Define your handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	// Create a middleware chain with standard middleware
	middleware := Chain(
		Recovery(),
		Logger(),
		CORS(),
		RateLimit(100),
	)

	// Apply the middleware chain to your handler
	http.Handle("/", middleware(handler))
	
	// Or use the StandardStack helper
	http.Handle("/api/", StandardStack()(handler))
	
	// Start the server
	http.ListenAndServe(":8080", nil)
}

func CustomizedExample() {
	// Customize logger
	loggerConfig := DefaultLoggerConfig()
	loggerConfig.ExcludePaths = append(loggerConfig.ExcludePaths, "/api/status")
	
	// Customize CORS
	corsConfig := DefaultCORSConfig()
	corsConfig.AllowOrigins = []string{"https://example.com", "https://api.example.com"}
	
	// Create middleware chain with custom config
	middleware := Chain(
		RecoveryWithConfig(DefaultRecoveryConfig()),
		LoggerWithConfig(loggerConfig),
		CORSWithConfig(corsConfig),
		BasicAuth("admin", "secure-password"),
	)
	
	http.Handle("/", middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Secured endpoint"))
	})))
	
	http.ListenAndServe(":8080", nil)
}
*/
