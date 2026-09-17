package main

import "fmt"

func main() {

	age := 70
	if age < 18 {
		fmt.Println("You are a minor.")
	} else if age >= 18 && age < 65 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a senior citizen.")
	}

	// if-else scope example


	if score:=50; score >= 90 { // if-else scope example
		fmt.Println("You got an A!")
	} else if score >= 80 {
		fmt.Println("You got a B!")
	} else if score >= 70 {
		fmt.Println("You got a C!")
	} else if score >= 60 {
		fmt.Println("You got a D!")
	} else {
		fmt.Println("You failed.")
	}


}