package main

import "fmt"

type CustomError struct {
	message string
	code    int
}

func (cu *CustomError) Error() string {
	return cu.message
}

func login(password string) error {

	//if not queal
	if password != "123" {
		return &CustomError{
			message: "password is not correct",
			code:    400,
		}
	}

	return nil

}

func main() {

	err := login("123")

	if err != nil {
		fmt.Println("Error", err)
	}

}
