package main

import (
	"errors"  // Provides error creation and manipulation functions
	"fmt"     // Provides formatting functions including fmt.Errorf for error wrapping
)

// read simulates accessing a file and returns an error if the filename is empty
// This creates our "base" or "root" error that will be wrapped later
func read(filename string) error {
	if filename == "" {
		// Create a simple error with a message
		return errors.New("the file is empty ")
	}
	return nil // No error occurred
}

// file is a higher-level function that calls read and wraps any returned error
// This demonstrates error wrapping - adding context to an error while preserving the original
func file() (string, error) {
	// Try to read a file with an empty name (which will cause an error)
	er := read("")
	
	// If read() returned an error, wrap it with additional context
	if er != nil {
		// The %w verb is special in fmt.Errorf - it wraps the error
		// This preserves the original error while adding new context
		// The wrapped error can later be extracted using errors.Unwrap()
		return "", fmt.Errorf("the lol lol \\ %w", er)
		// Note: Use %v instead of %w if you want to include the error message
		// but don't need to preserve the original error for unwrapping
	}
	
	// Return success if no error
	return "all good", nil
}

func main() {
	// Call file() and ignore the success return value, only capture the error
	_, abcd := file()
	
	// Check if an error occurred
	if abcd != nil {
		// Print the full error with its context
		fmt.Println("wow", abcd)
		
		// errors.Unwrap() extracts the original error that was wrapped with %w
		// This lets us access the underlying error without the added context
		// It's useful when you need to check the original error type or message
		unw := errors.Unwrap(abcd)
		fmt.Println("unwrapped me :", unw)
		
		// Note: You can chain Unwrap calls if there are multiple layers of wrapping
		// e.g., if we had another level: originalErr := errors.Unwrap(errors.Unwrap(abcd))
		
		// Instead of manually unwrapping, modern Go code typically uses:
		// - errors.Is(): To check if an error or any error it wraps matches a specific error
		// - errors.As(): To check if an error or any error it wraps is of a specific type
	} else {
		fmt.Println("mm", abcd)
	}
	
	// Common error unwrapping patterns:
	// 1. Sentinel error checking:
	//    if errors.Is(err, io.EOF) { /* handle EOF */ }
	//
	// 2. Type-based error handling:
	//    var netErr *net.OpError
	//    if errors.As(err, &netErr) { /* handle network error */ }
	//
	// 3. Custom error types with Unwrap method:
	//    type MyError struct { Err error, Context string }
	//    func (e *MyError) Unwrap() error { return e.Err }
	//    func (e *MyError) Error() string { return e.Context + ": " + e.Err.Error() }
}
