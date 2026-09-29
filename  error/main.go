package main

import (
	"errors"
	"fmt"
)

func main() {

	result, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

}

func divide(a int, b int) (int, error) {

	// how to use error in go
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
}
