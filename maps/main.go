// package main defines the entry package name for an executable Go program.
package main

// import "fmt" brings in Go's standard I/O formatting package used for printing output.
import "fmt"

// main is the entry point function where execution begins.
func main() {
	// make(map[string]int) initializes an empty hash map where keys are strings and values are integers.
	// We use the short variable declaration operator (:=) to assign it to 'myMap'.
	myMap := make(map[string]int)

	// Inserts or updates key "Alice" with the integer value 30.
	myMap["Alice"] = 30
	// Inserts or updates key "Bob" with the integer value 25.
	myMap["Bob"] = 25

	// Accesses the value for key "Alice" (returns 30) and prints it to the console.
	fmt.Println("Age of Alice:", myMap["Alice"])
	// Accesses the value for key "Bob" (returns 25) and prints it to the console.
	fmt.Println("Age of Bob:", myMap["Bob"])

	// The "comma-ok" idiom: retrieves the value into 'age' and a boolean flag into 'exists'.
	// If "Charlie" is in the map, exists is true; if not, exists is false and age gets the zero-value (0).
	age, exists := myMap["Charlie"]
	// Checks if the key was found in the map.
	if exists {
		// Runs if exists == true, printing Charlie's age.
		fmt.Println("Age of Charlie:", age)
	} else {
		// Runs if exists == false, letting us know the key was not found.
		fmt.Println("Charlie not found in the map.")
	}

	// delete() is a built-in function that removes the key "Bob" and its value from myMap.
	delete(myMap, "Bob")
	// Prints the updated state of myMap showing that "Bob" has been removed.
	fmt.Println("After deleting Bob:", myMap)
}