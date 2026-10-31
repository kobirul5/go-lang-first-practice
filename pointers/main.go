package main

import "fmt"

// ============================================================================
// GO POINTERS: COMPREHENSIVE GUIDE & PRACTICAL EXAMPLES
// ============================================================================
//
// Key Concepts:
// 1. '&' (Address-of Operator): Returns the memory address of a variable.
// 2. '*' (Dereference Operator):
//    - In types (e.g., *int, *string): Denotes a pointer type.
//    - In expressions (e.g., *p): Accesses or modifies the value stored at that address.
// 3. Zero Value: The zero value of any pointer in Go is `nil`.
// 4. Go does NOT support pointer arithmetic (unlike C/C++), making it memory-safe!
// ============================================================================

// ----------------------------------------------------------------------------
// EXAMPLE 1: Basic Pointer Operations (& and *)
// ----------------------------------------------------------------------------
func demoBasicPointers() {
	fmt.Println("--- Example 1: Basic Pointer Operations ---")

	a := 42
	p := &a // 'p' stores the memory address of 'a' (type: *int)

	fmt.Printf("Value of a      : %d\n", a)
	fmt.Printf("Memory addr of a: %p\n", &a)
	fmt.Printf("Value in p (addr): %p\n", p)
	fmt.Printf("Value via *p    : %d\n", *p)

	// Modifying 'a' through pointer 'p'
	*p = 100
	fmt.Printf("After *p = 100 -> Value of a: %d\n\n", a)
}

// ----------------------------------------------------------------------------
// EXAMPLE 2: Function Parameters (Pass-by-Value vs Pass-by-Pointer)
// ----------------------------------------------------------------------------
// In Go, everything is passed by value (a copy is created).
// Passing a pointer copies the memory address, allowing the original value to be changed.

func modifyValue(val int) {
	val = 999 // Only changes the local copy inside this function
}

func modifyPointer(val *int) {
	*val = 999 // Mutates the original variable at that memory address
}

func demoFunctionArguments() {
	fmt.Println("--- Example 2: Pass by Value vs Pass by Pointer ---")

	x := 10
	fmt.Printf("Initial x: %d\n", x)

	modifyValue(x)
	fmt.Printf("After modifyValue(x)  : %d (No change! Copy was modified)\n", x)

	modifyPointer(&x)
	fmt.Printf("After modifyPointer(&x): %d (Updated! Modified via memory address)\n\n", x)
}

// ----------------------------------------------------------------------------
// EXAMPLE 3: Pointers with Structs (Value Receiver vs Pointer Receiver)
// ----------------------------------------------------------------------------
type Person struct {
	Name string
	Age  int
}

// Value receiver: operates on a COPY of Person
func (p Person) HaveBirthdayValue() {
	p.Age++ // Will NOT persist outside this method
}

// Pointer receiver: operates on the ORIGINAL Person
func (p *Person) HaveBirthdayPointer() {
	// In Go, p.Age is syntactic sugar for (*p).Age (automatic dereferencing)
	p.Age++ // Updates the original struct!
}

func demoStructPointers() {
	fmt.Println("--- Example 3: Structs & Pointer Receivers ---")

	user := Person{Name: "Rahim", Age: 25}
	fmt.Printf("Initial user: %+v\n", user)

	user.HaveBirthdayValue()
	fmt.Printf("After HaveBirthdayValue()  : %+v (Age did not change)\n", user)

	user.HaveBirthdayPointer()
	fmt.Printf("After HaveBirthdayPointer(): %+v (Age updated successfully!)\n\n", user)
}

// ----------------------------------------------------------------------------
// EXAMPLE 4: The new() Built-in Function
// ----------------------------------------------------------------------------
// new(T) allocates zeroed memory for type T and returns a pointer (*T).
func demoNewFunction() {
	fmt.Println("--- Example 4: The new() Built-in Function ---")

	// ptr is of type *int, initialized with memory pointing to 0
	ptr := new(int)

	fmt.Printf("Address in ptr : %p\n", ptr)
	fmt.Printf("Default value  : %d (Zero value of int)\n", *ptr)

	*ptr = 55
	fmt.Printf("Updated value  : %d\n\n", *ptr)
}

// ----------------------------------------------------------------------------
// EXAMPLE 5: Nil Pointers & Safe Defensive Checks
// ----------------------------------------------------------------------------
// Uninitialized pointers hold 'nil'. Dereferencing a nil pointer causes panic!
func demoNilPointers() {
	fmt.Println("--- Example 5: Nil Pointer Checks (Preventing Panics) ---")

	var safePtr *string // safePtr is nil by default

	if safePtr == nil {
		fmt.Println("safePtr is nil! Checking before dereferencing prevents panic.")
	}

	msg := "Golang is awesome"
	safePtr = &msg

	if safePtr != nil {
		fmt.Printf("Now safe to read: %s\n\n", *safePtr)
	}
}

// ----------------------------------------------------------------------------
// EXAMPLE 6: Pointer to Pointer (Double Pointer)
// ----------------------------------------------------------------------------
// Useful for updating where an existing pointer points.
func demoPointerToPointer() {
	fmt.Println("--- Example 6: Pointer to Pointer (**int) ---")

	val := 500
	p1 := &val // p1 is of type *int
	p2 := &p1  // p2 is of type **int (stores the address of pointer p1)

	fmt.Printf("val        : %d\n", val)
	fmt.Printf("*p1        : %d\n", *p1)
	fmt.Printf("**p2       : %d (dereferenced twice to reach val)\n", **p2)

	// Updating value via double pointer
	**p2 = 750
	fmt.Printf("After **p2 = 750 -> val is now: %d\n\n", val)
}

// ----------------------------------------------------------------------------
// EXAMPLE 7: Real-world Pattern: Optional / Nullable Fields in Structs
// ----------------------------------------------------------------------------
// In APIs or DB models, pointers distinguish between:
// - zero/empty value (e.g., 0, "")
// - absence of a value (nil / NULL)
type Profile struct {
	Username string
	Bio      *string // Optional field: nil means bio was not provided
	Age      *int    // Optional field: nil means age was not set (distinct from 0)
}

func demoOptionalFields() {
	fmt.Println("--- Example 7: Real-world Use Case (Optional/Nullable Fields) ---")

	// User 1: did not provide bio or age
	user1 := Profile{
		Username: "john_doe",
		Bio:      nil,
		Age:      nil,
	}

	// User 2: provided bio and age
	bioText := "Backend Software Engineer"
	ageVal := 28
	user2 := Profile{
		Username: "alice_dev",
		Bio:      &bioText,
		Age:      &ageVal,
	}

	printProfile(user1)
	printProfile(user2)
}

func printProfile(p Profile) {
	fmt.Printf("Username: %s\n", p.Username)

	if p.Bio != nil {
		fmt.Printf("  Bio: %s\n", *p.Bio)
	} else {
		fmt.Println("  Bio: [Not provided]")
	}

	if p.Age != nil {
		fmt.Printf("  Age: %d\n", *p.Age)
	} else {
		fmt.Println("  Age: [Not provided]")
	}
}

// ----------------------------------------------------------------------------
// MAIN FUNCTION
// ----------------------------------------------------------------------------
func main() {
	fmt.Println("==================================================")
	fmt.Println("          GO POINTERS MASTERCLASS DEMO            ")
	fmt.Println("==================================================")
	fmt.Println()

	demoBasicPointers()
	demoFunctionArguments()
	demoStructPointers()
	demoNewFunction()
	demoNilPointers()
	demoPointerToPointer()
	demoOptionalFields()

	fmt.Println("==================================================")
	fmt.Println("All pointer examples executed successfully!")
	fmt.Println("==================================================")
}