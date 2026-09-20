package main

import "fmt"

func main() {
	// slice
	var orders = [6]int{1, 2, 3, 4, 5, 6}
	slice := orders[:]

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice)
	}

	// slice capacity
	fmt.Println("Slice capacity:", cap(slice))
	// slice length
	fmt.Println("Slice length:", len(slice))

	// what is slice ? slice length and capacity?
	//answer: slice is a reference type, it is a view of an array. Slice length is the number of elements in the slice, and slice capacity is the number of elements in the underlying array starting from the first element in the slice.
}