package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
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


	var person2 Person
	jsonData := `{"name":"Jane Smith","age":25,"city":"Los Angeles"}`

	error := json.Unmarshal([]byte(jsonData), &person2)

	if error != nil {
		fmt.Println("Error unmarshalling JSON:", error)
	}

	fmt.Println(person2)

}
