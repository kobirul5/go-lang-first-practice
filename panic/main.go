package main

import (
	"fmt"
	"log"
	"sync"
)

// ============================================================================
// GO ERROR HANDLING: PANIC, RECOVER, LOG.FATAL & COMMON PANIC EXAMPLES
// ============================================================================
//
// In Go, normal errors are returned as values (e.g., `return val, err`).
// Panic is reserved for truly unrecoverable or unexpected runtime anomalies.
//
// 1. panic():
//    - Stops normal execution flow of the current goroutine immediately.
//    - Stack Unwinding: executes all deferred functions (`defer`) up the call stack.
//    - If not caught by `recover()`, prints the panic message, stack trace, and crashes.
//
// 2. recover():
//    - Must be called directly inside a `defer` function.
//    - Stops unwinding and captures the panic value (type `any`).
//    - Returns `nil` if there is no active panic.
//
// 3. Goroutine Boundary (Important!):
//    - Panics DO NOT cross goroutines! A panic in a child goroutine MUST be
//      recovered inside that same goroutine, otherwise the whole program crashes.
// ============================================================================

// ----------------------------------------------------------------------------
// EXAMPLE 1: Basic Panic & Recover (Manual Panic)
// ----------------------------------------------------------------------------
func demoBasicPanic() {
	fmt.Println("\n[Demo 1] 1. Starting execution...")

	defer func() {
		fmt.Println("[Demo 1] 3. Deferred cleanup function is running...")
		if r := recover(); r != nil {
			fmt.Printf("[Demo 1] 4. Successfully recovered! Panic Reason: %v\n", r)
		}
	}()

	fmt.Println("[Demo 1] 2. Triggering explicit panic()...")
	panic("Custom failure: Database connection timeout!")
}

// ----------------------------------------------------------------------------
// EXAMPLE 2: Common Runtime Panics (স্বাভাবিক যেসব কারণে Go-তে Panic হয়)
// ----------------------------------------------------------------------------

// 2a. Nil Pointer Dereference
func demoNilPointer() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("[Demo 2a: Nil Pointer] Caught panic gracefully: %v\n", r)
		}
	}()

	var ptr *string // ptr is nil
	// Dereferencing a nil pointer causes a runtime panic:
	// "runtime error: invalid memory address or nil pointer dereference"
	*ptr = "Hello Go"
}

// 2b. Slice Index Out of Range
func demoSliceOutOfBounds() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("[Demo 2b: Index Out of Bounds] Caught panic: %v\n", r)
		}
	}()

	numbers := []int{10, 20, 30} // length is 3 (indices: 0, 1, 2)
	// Accessing index 5 triggers: "runtime error: index out of range [5] with length 3"
	_ = numbers[5]
}

// 2c. Panics with Channels (Closing a closed channel or sending on closed channel)
func demoChannelPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("[Demo 2c: Channel Panic] Caught panic: %v\n", r)
		}
	}()

	ch := make(chan int, 1)
	close(ch)

	// Sending to a closed channel causes panic: "send on closed channel"
	ch <- 42

	// Similarly, closing an already closed channel causes: "close of closed channel"
	// close(ch)
}

// 2d. Type Assertion Failure without comma-ok idiom
func demoTypeAssertionPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("[Demo 2d: Type Assertion] Caught panic: %v\n", r)
		}
	}()

	var data any = "Hello, World!" // dynamic type is string

	// Direct assertion to int (without ok) panics:
	// "interface conversion: interface {} is string, not int"
	num := data.(int)
	fmt.Println(num)
}

// ----------------------------------------------------------------------------
// EXAMPLE 3: Panic to Error Conversion (Named Return Values)
// ----------------------------------------------------------------------------
// Real-world pattern: converting unexpected panics into clean Go errors.
func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("division failed: %v", r)
		}
	}()

	// Integer divide by zero triggers a runtime panic: "integer divide by zero"
	result = a / b
	return result, nil
}

// ----------------------------------------------------------------------------
// EXAMPLE 4: Panics Across Goroutines (Goroutine Boundary Trap)
// ----------------------------------------------------------------------------
// WARNING: A defer/recover in main() will NOT catch a panic from another goroutine!
// Every goroutine must handle its own panic.
func demoGoroutinePanic() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		// CRITICAL: recover() must be inside THIS goroutine
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("[Demo 4: Goroutine] Caught panic inside background worker: %v\n", r)
			}
		}()

		fmt.Println("[Demo 4: Goroutine] Background worker started...")
		panic("Worker encountered unexpected failure!")
	}()

	wg.Wait()
	fmt.Println("[Demo 4: Goroutine] Goroutine completed safely without crashing main!")
}

// ----------------------------------------------------------------------------
// EXAMPLE 5: Re-panicking (Rethrowing Panic)
// ----------------------------------------------------------------------------
// Sometimes we only want to recover specific expected issues.
// If it's something unexpected or severe, we re-throw it using panic(r).
func demoRepanic(shouldRethrow bool) {
	defer func() {
		r := recover()
		if r == nil {
			return
		}

		if shouldRethrow {
			fmt.Printf("[Demo 5: Re-panic] Critical panic detected (%v), re-throwing...\n", r)
			// Re-throw panic so upper layers or crash reporter can see it
			panic(r)
		}

		fmt.Printf("[Demo 5: Re-panic] Recovered successfully from: %v\n", r)
	}()

	panic("Database disk corrupted!")
}

// ----------------------------------------------------------------------------
// EXAMPLE 6: log.Panic vs log.Fatal
// ----------------------------------------------------------------------------
// - log.Panic(): Prints log, triggers panic(), and allows defers to run (recoverable).
// - log.Fatal(): Prints log, calls os.Exit(1), SKIPS all defers (unrecoverable).
func demoLogPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("[Demo 6: log.Panic] Successfully caught log.Panic! Reason: %v\n", r)
		}
	}()

	fmt.Println("[Demo 6: log.Panic] Calling log.Panic (defers will execute)...")
	log.Panic("Critical configuration missing! (Triggered via log.Panic)")
}

func main() {
	fmt.Println("==================================================")
	fmt.Println("       GO PANIC & RECOVER ADVANCED EXAMPLES       ")
	fmt.Println("==================================================")

	// 1. Basic Panic & Recover
	fmt.Println("\n--- Scenario 1: Manual Panic & Recover ---")
	demoBasicPanic()

	// 2. Common Runtime Panics
	fmt.Println("\n--- Scenario 2: Common Runtime Panics ---")
	demoNilPointer()
	demoSliceOutOfBounds()
	demoChannelPanic()
	demoTypeAssertionPanic()

	// 3. Converting Panic to clean error
	fmt.Println("\n--- Scenario 3: Converting Panic to Error ---")
	res, err := safeDivide(10, 0)
	if err != nil {
		fmt.Printf("[main] Caught error: %v\n", err)
	} else {
		fmt.Printf("[main] Result: %d\n", res)
	}

	// 4. Goroutine Safe Panic Recovery
	fmt.Println("\n--- Scenario 4: Panic Inside Goroutines ---")
	demoGoroutinePanic()

	// 5. Re-panicking pattern
	fmt.Println("\n--- Scenario 5: Re-panicking ---")
	demoRepanic(false) // caught safely without re-throwing

	// 6. log.Panic with recovery
	fmt.Println("\n--- Scenario 6: log.Panic vs log.Fatal ---")
	demoLogPanic()

	fmt.Println("\n==================================================")
	fmt.Println("[main] All panic examples completed without crash!")
	fmt.Println("==================================================")
}