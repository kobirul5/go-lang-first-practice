package main

import "fmt"

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

func main() {
	fmt.Println("Inside main: calling doSomting()...")

	// Calling the function that triggers the panic
	doSomting()

	// This line will NEVER be reached because the program crashes during doSomting()
	fmt.Println("Inside main: completed successfully.")
}