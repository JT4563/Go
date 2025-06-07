package main

import "fmt"

// Request represents an HTTP request
type Request struct {
    Path   string
    Method string
}

// Middleware interface defines a contract for request handlers
type Middleware interface {
    Handle(Request) Request
}

// Logger is a middleware that logs requests
type Logger struct{}

func (l Logger) Handle(req Request) Request {
    fmt.Println("Logging request:", req.Method, req.Path)
    return req
}

// Auth is a middleware that authenticates requests
type Auth struct{}

func (a Auth) Handle(req Request) Request {
    fmt.Println("Checking auth for request to:", req.Path)
    return req
}

func main() {
    // Create a sample request
    request := Request{
        Path:   "/api/users",
        Method: "GET",
    }
    
    fmt.Println("Initial request:", request)
    
    // Create middleware chain
    middlewares := []Middleware{Logger{}, Auth{}}
    
    // Apply each middleware in sequence
    for _, m := range middlewares {
        request = m.Handle(request)
    }
    
    fmt.Println("Request processing complete")
}