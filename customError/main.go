package main

import (
	"fmt"
)

// CustomError struct defines our custom error type with additional fields.
// It contains a message string and an HTTP-like status code.
type CustomError struct {
	message string
	code    int
}

// Error implements the built-in error interface for CustomError.
// By adding this method, CustomError can be used wherever an error is expected.
func (ce *CustomError) Error() string {
	return fmt.Sprintf("Error Code %d: %s", ce.code, ce.message)
}

// login simulates a login function that returns an error if authentication fails.
func login(password string) error {
	// If password doesn't match, return our custom error
	if password != "secret123" {
		return &CustomError{
			message: "Incorrect password provided",
			code:    401, // Unauthorized
		}
	}

	// Return nil if there's no error
	return nil
}

func main() {
	// Example 1: Unsuccessful login (Will throw custom error)
	fmt.Println("--- Example 1: Invalid Password ---")
	err := login("wrong_password")
	
	if err != nil {
		// We can print the error directly, which calls our Error() method
		fmt.Println("General Error:", err)

		// Type assertion: We can extract the specific CustomError to access its fields
		if customErr, ok := err.(*CustomError); ok {
			fmt.Printf("Extracted Code: %d\n", customErr.code)
			fmt.Printf("Extracted Message: %s\n", customErr.message)
		}
	}

	// Example 2: Successful login
	fmt.Println("\n--- Example 2: Valid Password ---")
	err2 := login("secret123")
	if err2 == nil {
		fmt.Println("Login successful! No error.")
	}
}
