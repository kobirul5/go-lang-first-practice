package main

import "fmt"

func deferredFunction(result int) {

	fmt.Println("I am from deferred function and result is ", result)

}

func exampleDefer() int{

	result := 10
	defer deferredFunction(result)
	fmt.Println("I am from exampleDefer function and result is ", result)

	// multiply deferred 
	defer fmt.Println("I am from exampleDefer function and result is ",1)
	defer fmt.Println("I am from exampleDefer function and result is ",2)
	defer fmt.Println("I am from exampleDefer function and result is ",3)
	defer fmt.Println("I am from exampleDefer function and result is ",4)
	defer fmt.Println("I am from exampleDefer function and result is ",5)
	result += 20

	return result 

}

func namedDefer() (result int) {
	result = 10

 defer func() {
	result += 20
        fmt.Println("I am from deferred function and result is", result)
    }()


	fmt.Println("I am from namedDefer function and result is ", result)
	result += 20
	return
}

func main() {
	defer fmt.Println("I am from Deferred print call ")
	fmt.Println("I am from main function")


	exampleDefer()
	namedDefer()

}
