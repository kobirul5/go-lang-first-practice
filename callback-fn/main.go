package main

import "fmt"

func calculate(a int, b int, callback func(x int, y int) int) int {
	return callback(a, b)
}


func main() {


	add := func(x int, y int) int {
		return x + y
	}

	subtract := func(x int, y int) int {
		return x - y
	}

	multiply := func(x int, y int) int {
		return x * y
	}

	divide := func(x int, y int) int {
		return x / y
	}

	calc1 := calculate(10, 5, add)
	calc2 := calculate(10, 5, subtract)
	calc3 := calculate(10, 5, multiply)
	calc4 := calculate(10, 5, divide)

	fmt.Println("Addition:", calc1)
	fmt.Println("Subtraction:", calc2)
	fmt.Println("Multiplication:", calc3)
	fmt.Println("Division:", calc4)


	

}