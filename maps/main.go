package main

import "fmt"

// maps
func main() {


	myMap := make(map[string]int)

	myMap["Alice"] = 30
	myMap["Bob"] = 25

	fmt.Println("Age of Alice:", myMap["Alice"])
	fmt.Println("Age of Bob:", myMap["Bob"])

	// Checking if a key exists
	age, exists := myMap["Charlie"]
	if exists {
		fmt.Println("Age of Charlie:", age)
	} else {
		fmt.Println("Charlie not found in the map.")
	}

	// Deleting a key-value pair
	delete(myMap, "Bob")
	fmt.Println("After deleting Bob:", myMap)

}