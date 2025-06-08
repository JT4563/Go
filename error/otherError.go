package main

import (
	"errors"
	"fmt"
)

// Fuer is a sentinel error that represents a specific error condition
// By declaring this as a package-level variable, we can check for this exact error later
var Fuer = errors.New("the file is empty")

// read simulates a function that might return different types of errors
// It takes a filename string and returns an error if something goes wrong
func read(filename string) error {
	// Check if the filename is empty
	if filename == "" {
		// Return our predefined sentinel error when the filename is empty
		return Fuer 
	}
	// Return nil to indicate success (no error)
	return nil
}

func main() {
	// Call the read function with an empty string, which will trigger our error
	er := read("")

	// Error handling pattern in Go:
	// 1. First, check for specific errors using errors.Is()
	//errors.Is check kore je error ta direct oi ErrFileEmpty error kina, ba oi error ke wrap kora ache kina.
	//    errors.Is() safely compares errors even when they're wrapped in other errors
	if errors.Is(er, Fuer) {
		fmt.Println("the file is empty")
	} 
	// 2. Then, check for any other error conditions
	//    This ensures we handle unexpected errors too
	else if er != nil {
		fmt.Println("other error:", er)
	} 
	// 3. Finally, handle the success case
	//    When er == nil, it means no error occurred
	else {
		fmt.Println("no error", er)
	}

	// Note: Other error checking patterns in Go include:
	// - errors.As(): To check if an error is of a specific custom type
	// - Type assertions: To access fields of custom error types
	// - Custom error types: For richer error information
}
