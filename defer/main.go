package main

import "fmt"

// deferredFunction is a simple function that will be called using defer later.
func deferredFunction(result int) {
	fmt.Printf("-> Inside deferredFunction: received result as %d\n", result)
}

// exampleDefer demonstrates how defer evaluates arguments immediately,
// and how multiple defers are executed in LIFO (Last-In-First-Out) order.
func exampleDefer() int {
	fmt.Println("\n--- Running exampleDefer ---")
	result := 10
	
	// 'result' is evaluated as 10 right now, even though the function 
	// call is deferred until exampleDefer returns.
	defer deferredFunction(result)

	fmt.Printf("Inside exampleDefer (Before addition): result is %d\n", result)

	// Multiple defers execute in LIFO (Last In, First Out) order.
	// So it will print 3, then 2, then 1.
	defer fmt.Println("Deferred LIFO Test: 1")
	defer fmt.Println("Deferred LIFO Test: 2")
	defer fmt.Println("Deferred LIFO Test: 3")
	
	result += 20
	fmt.Printf("Inside exampleDefer (After addition): result is %d\n", result)

	return result
}

// namedDefer demonstrates how defer can modify named return values.
func namedDefer() (result int) {
	fmt.Println("\n--- Running namedDefer ---")
	result = 10

	// This anonymous function is deferred. 
	// Because 'result' is a named return variable, this deferred function 
	// can actually read and change its final value right before the function returns.
	defer func() {
		result += 20
		fmt.Printf("-> Inside named defer closure: result updated to %d\n", result)
	}()

	fmt.Printf("Inside namedDefer: result is %d\n", result)
	result += 20 // result becomes 30
	
	// When return is called, the deferred function runs and adds 20 more to result (makes it 50)
	return
}

func main() {
	// This defer will run at the very end of the main function, right before the program exits.
	defer fmt.Println("\n[Main Defer] -> Program is about to exit. Goodbye!")
	
	fmt.Println("[Main Function] -> Program started.")

	// Call the examples
	exampleDefer()
	
	finalNamedResult := namedDefer()
	fmt.Printf("Result returned from namedDefer: %d\n", finalNamedResult)
}
