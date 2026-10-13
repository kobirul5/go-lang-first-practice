package main

import (
	"fmt"
	"log"
)

// ============================================================================
// GO ERROR HANDLING: PANIC, RECOVER, LOG.FATAL & OS.EXIT EXPLAINED
// ============================================================================
//
// In Go, normal errors are returned as values (e.g., `return val, err`).
// However, Go also provides panic, recover, and process exit mechanisms.
//
// 1. panic():
//    - Stops normal execution flow of the current goroutine immediately.
//    - Stack Unwinding: runs all deferred (`defer`) functions in the current function,
//      returns to the caller, runs its defers, and continues all the way up the call stack.
//    - If not caught by `recover()`, prints the panic message, full stack trace, and crashes the program.
//    - Goroutine Scope: Panic cannot cross goroutine boundaries. A panic in one goroutine
//      must be recovered in that same goroutine, otherwise the whole application crashes.
//
// 2. recover():
//    - Built-in function used to regain control of a panicking goroutine.
//    - CRITICAL RULE: `recover()` is ONLY effective directly inside a deferred function (`defer func() { ... }`).
//      Calling `recover()` outside a deferred function always returns `nil` and has no effect.
//    - If the current goroutine is panicking, `recover()` stops the panic and returns the value
//      passed to `panic(...)` (type `any` / `interface{}`).
//    - If there is no active panic, `recover()` simply returns `nil`.
//
// 3. log.Fatal() vs log.Panic() vs os.Exit():
//    - log.Fatal(): Prints the message with a timestamp and calls `os.Exit(1)`.
//      --> WARNING: Defers DO NOT run! Program exits immediately.
//    - log.Panic(): Prints the message with a timestamp and calls `panic()`.
//      --> Defers DO run! It can be caught with `recover()`.
//    - os.Exit(code): Terminates the process immediately with the given status code.
//      --> No defers, no stack trace.
// ============================================================================

// doSomting demonstrates how panic and recover interact with defer.
// It also shows how a deferred function can inspect and handle the panic value.
func doSomting() {
	fmt.Println("\n[doSomting] 1. Starting execution...")

	// ------------------------------------------------------------------------
	// DEFER & RECOVER BLOCK:
	// 'defer' schedules this anonymous function to run when 'doSomting' exits,
	// whether it exits normally OR abnormally (due to a panic).
	// ------------------------------------------------------------------------
	defer func() {
		fmt.Println("[doSomting] 3. Deferred cleanup function is now running...")

		// recover() checks if there is an active panic in this goroutine.
		// If yes, it captures the panic value and stops the stack unwinding.
		r := recover()
		if r != nil {
			// 'r' contains the value passed to panic() — in this case: "Something went wrong!"
			fmt.Printf("[doSomting] 4. Successfully recovered from panic! Reason: %v\n", r)
			fmt.Println("[doSomting] 5. Goroutine is back to normal. The caller continues execution!")
		} else {
			// If doSomting ended normally without any panic, r would be nil.
			fmt.Println("[doSomting] No panic detected.")
		}
	}()

	fmt.Println("[doSomting] 2. About to trigger panic...")

	// Triggering panic:
	// Once this line executes:
	// - Line 57 onward in this function will be SKIPPED.
	// - Go immediately enters panicking mode.
	// - Go searches for deferred functions in this scope (the defer above runs next!).
	panic("Something went wrong!")

	// ------------------------------------------------------------------------
	// UNREACHABLE CODE:
	// Any statements placed directly after panic() in the same block will never run.
	// ------------------------------------------------------------------------
	// fmt.Println("[doSomting] This line will never execute.")
}

// safeDivide demonstrates a practical real-world pattern:
// Converting a panic (e.g. divide by zero) into a clean error using named return values and recover().
func safeDivide(a, b int) (result int, err error) {
	// A deferred function can read and modify named return values (result, err)
	defer func() {
		if r := recover(); r != nil {
			// Catch any runtime panic and convert it into a standard Go error
			err = fmt.Errorf("runtime panic caught: %v", r)
		}
	}()

	// In Go, integer division by zero triggers a runtime panic: "integer divide by zero"
	result = a / b
	return result, nil
}

// demonstrateFatal shows how log.Fatal operates and why it differs fundamentally from panic.
func demonstrateFatal() {
	// Notice this deferred function:
	defer func() {
		// CRITICAL GOTCHA:
		// This line will NEVER print!
		// log.Fatal() calls os.Exit(1) directly, terminating the OS process immediately
		// without unwinding the stack or executing any pending deferred calls.
		fmt.Println("[demonstrateFatal] This defer will NEVER execute!")
	}()

	fmt.Println("\n[demonstrateFatal] About to call log.Fatal...")
	log.Fatal("Fatal error occurred! Process exits immediately (defers are skipped).")
}

func main() {
	fmt.Println("==================================================")
	fmt.Println("       GO PANIC, RECOVER & LOG.FATAL DEMO         ")
	fmt.Println("==================================================")

	// ------------------------------------------------------------------------
	// SCENARIO 1: Basic Panic & Recover
	// Shows that recover stops panic from crashing the program, allowing main()
	// to continue executing normally.
	// ------------------------------------------------------------------------
	fmt.Println("\n--- Scenario 1: Panic with Recover ---")
	doSomting()
	fmt.Println("[main] Execution continued successfully after doSomting()!")

	// ------------------------------------------------------------------------
	// SCENARIO 2: Panic-to-Error Conversion (Idiomatic Pattern)
	// Often used in libraries or HTTP recovery middleware to prevent one bad
	// request from crashing the entire server.
	// ------------------------------------------------------------------------
	fmt.Println("\n--- Scenario 2: Handling Panic with Named Return ---")
	res, err := safeDivide(10, 0)
	if err != nil {
		fmt.Printf("[main] Caught division error gracefully: %v\n", err)
	} else {
		fmt.Printf("[main] Division result: %d\n", res)
	}

	// ------------------------------------------------------------------------
	// SCENARIO 3: log.Fatal()
	// UNCOMMENT the lines below to observe log.Fatal() in action.
	// Keep in mind: because log.Fatal() calls os.Exit(1), the program will
	// terminate immediately and any lines below it will not run!
	// ------------------------------------------------------------------------
	// fmt.Println("\n--- Scenario 3: Testing log.Fatal ---")
	// demonstrateFatal()

	fmt.Println("\n==================================================")
	fmt.Println("[main] All demonstrations completed successfully.")
	fmt.Println("==================================================")
}