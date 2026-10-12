package main

import (
	"fmt"
	"log"
)

// doSomting demonstrates how panic works in Go.
// 'panic' is a built-in function that stops the normal execution flow of the current goroutine.
// When a function calls panic:
// 1. Normal execution stops immediately.
// 2. Any deferred functions (using 'defer') in that function are executed.
// 3. Control returns to the caller, which also begins panicking and running its deferred calls.
// 4. This unwinding continues up the call stack until the program terminates and prints a stack trace (unless caught with recover()).
func doSomting() {
	fmt.Println("Inside doSomting: starting execution...")

	// Deferred functions always run when the surrounding function exits,
	// even if the exit is caused by a panic.
	defer func() {
		fmt.Println("Deferred function executed inside doSomting!")
		r := recover()
		if r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// Panic should generally be reserved for unrecoverable errors (e.g., critical misconfiguration, nil pointer).
	panic("Something went wrong!")

	// Unreachable code: anything after panic in the same block will NOT execute.
	// fmt.Println("Inside doSomting: this line will never run.")
}

// demonstrateFatal shows how log.Fatal works.
// Key difference from panic:
// - log.Fatal() prints the message with timestamp and immediately calls os.Exit(1).
// - CRITICAL: log.Fatal() DOES NOT run deferred (defer) functions!
func demonstrateFatal() {
	defer func() {
		// WARNING: This will NOT run when log.Fatal is called!
		fmt.Println("This defer will NEVER execute because log.Fatal calls os.Exit directly!")
	}()

	fmt.Println("Inside demonstrateFatal: about to call log.Fatal...")
	log.Fatal("Fatal error occurred! Program exits immediately.")
}

func main() {
	// 1. Uncomment this to see panic & recover in action:
	// fmt.Println("--- Testing Panic ---")
	// doSomting()

	// 2. Testing log.Fatal():
	fmt.Println("--- Testing log.Fatal ---")
	demonstrateFatal()

	// This line will never execute after log.Fatal
	fmt.Println("Inside main: completed successfully.")
}