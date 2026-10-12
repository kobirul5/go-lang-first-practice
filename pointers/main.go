package main

import "fmt"

func main() {
	// 1. Declare and initialize a standard integer variable 'a' with the value 42
	a := 42

	// 2. '&' (address-of) operator gets the memory address of variable 'a'.
	// 'p' is a pointer variable of type *int that stores the memory address of 'a'.
	p := &a

	// 3. Directly update the value of 'a' to 21
	a = 21

	fmt.Println(a)  // Output: 21 (current value of 'a')
	fmt.Println(p)  // Output: 0x... (memory address of 'a', e.g., 0xc000018088)
	fmt.Println(*p) // Output: 21 ('*' dereferences 'p' to read the value stored at that address)

	// 4. Modify the value via pointer dereferencing:
	// '*p = 42' goes to the memory address stored in 'p' (which is 'a') and sets its value to 42
	*p = 42

	fmt.Println(a)  // Output: 42 ('a' also changes because its underlying memory was updated via 'p')
	fmt.Println(p)  // Output: 0x... (memory address remains the same)
	fmt.Println(*p) // Output: 42 (reading the updated value through pointer dereferencing)
}