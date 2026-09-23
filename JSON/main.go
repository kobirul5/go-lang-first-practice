package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
	City string
}

func main() {
	// Create a new Person instance
	person := Person{
		Name: "John Doe",
		Age:  30,
		City: "New York",
	}

	rawJson, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
	}
	fmt.Println(string(rawJson))

}
