package main

import "fmt"

func main(){
	a := 42
	p := &a

	a = 21

	fmt.Println(a)
	fmt.Println(p)
	fmt.Println(*p)

	*p = 42

	fmt.Println(a)
	fmt.Println(p)
	fmt.Println(*p)
}