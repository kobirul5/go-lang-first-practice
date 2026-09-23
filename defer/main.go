package main

import "fmt"

func deferredFunction(result int) {

	fmt.Println("I am from deferred function and result is ", result)

}

func exampleDefer() {

	result := 10
	defer deferredFunction(result)
	fmt.Println("I am from exampleDefer function and result is ", result)

}

func main() {
	defer fmt.Println("I am from Deferred print call ")
	fmt.Println("I am from main function")

	exampleDefer()

}
