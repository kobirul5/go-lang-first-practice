package main

import "fmt"

//  isAdmin := true // wrong way, because short way not working out of main function


func makeCoffee(kind string){
	fmt.Printf("maing %s coffee...\n" , kind)
}

func main(){
	
	// fmt.Println("Hello, World!")

	// var name string = "Kobirul"
	// name:= "kobirul 2"  // short way , most used in real project

	//  grouped variables
	// var (
	// 	name string ="liton"
	// 	age int =5
	// )
	// fmt.Println(name, age)

	// multiple variable declaration
	// var x, y int
	//  x = 20
	//  y = 30
	//  fmt.Println(y-x)
	 
	//  var a, b string = "Ph", "hero"
	//  fmt.Println(a+b)


	// constant
	// const pi = 3.14
	//  fmt.Println(pi)


	// var age int
	// fmt.Println(age)//0


	// var name string
	// fmt.Println(name)//""

	// var isAdmin bool
	// fmt.Println(isAdmin)//false

	// var score float64
	// fmt.Println(score)//false

	//basic function
	makeCoffee("black")
	makeCoffee("cold")

}