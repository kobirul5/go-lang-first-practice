package main


import "fmt"

type additionalInformation struct {
	Address string
	Phone   string
}

type Person struct {
	Name string
	Age  int
	Info additionalInformation
}

// func main() {
// 	p1 := Person{Name: "Alice", Age: 30}
// 	p2 := Person{Name: "Bob", Age: 25}

// 	fmt.Println("Person 1:", p1)
// 	fmt.Println("Person 2:", p2)

// 	// Accessing fields
// 	fmt.Println("Name of Person 1:", p1.Name)
// 	fmt.Println("Age of Person 2:", p2.Age)

// 	// Modifying fields
// 	p1.Age = 31
// 	fmt.Println("Updated Age of Person 1:", p1.Age)
// }


// recieiver function

func (p Person) greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func main() {
	p1 := Person{Name: "Alice", Age: 30}
	p2 := Person{Name: "Bob", Age: 25}

	p1.greet()
	p2.greet()
}	