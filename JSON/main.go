package main // Declares this file as part of the main package, making it an executable program.

import ( // Starts the import block to include standard library packages.
	"encoding/json" // Imports the "encoding/json" package for encoding and decoding JSON.
	"fmt" // Imports the "fmt" package for formatted I/O, like printing to the console.
)

type Person struct { // Defines a new struct type named 'Person' to hold related data.
	Name string `json:"name"` // Declares a string field 'Name' and a struct tag for JSON encoding/decoding as "name".
	Age  int    `json:"age"`  // Declares an integer field 'Age' and a struct tag for JSON encoding/decoding as "age".
	City string `json:"city"` // Declares a string field 'City' and a struct tag for JSON encoding/decoding as "city".
}

func main() { // Defines the main function, which is the entry point of the executable program.
	// Create a new Person instance
	person := Person{ // Initializes a new variable 'person' of type 'Person' with specific values.
		Name: "John Doe", // Sets the 'Name' field of the struct to "John Doe".
		Age:  30,         // Sets the 'Age' field of the struct to 30.
		City: "New York", // Sets the 'City' field of the struct to "New York".
	}

	rawJson, err := json.Marshal(person) // Converts the 'person' struct into JSON byte slice. Returns the JSON bytes and an error if any.
	if err != nil { // Checks if an error occurred during the JSON marshaling process.
		fmt.Println("Error marshalling JSON:", err) // Prints the error message if marshaling failed.
	}
	fmt.Println(string(rawJson)) // Converts the raw JSON byte slice to a string and prints it to the console.


	var person2 Person // Declares a new variable 'person2' of type 'Person' with zero values (empty).
	jsonData := `{"name":"Jane Smith","age":25,"city":"Los Angeles"}` // Defines a string containing valid JSON data.

	error := json.Unmarshal([]byte(jsonData), &person2) // Decodes the JSON string into the 'person2' struct. We pass the address of 'person2' (&person2).

	if error != nil { // Checks if an error occurred during the JSON unmarshaling process.
		fmt.Println("Error unmarshalling JSON:", error) // Prints the error message if unmarshaling failed.
	}

	fmt.Println(person2) // Prints the 'person2' struct to see the populated values.
	// Output: // This is a comment indicating what output to expect from the previous print.
	// {Jane Smith 25 Los Angeles} // The expected output of printing 'person2'.
	// Output: // This seems to be an old or duplicate comment for expected output.
	// {"name":"John Doe","age":30,"city":"New York"} // The expected output of the first print statement (printing the raw JSON string).

}
