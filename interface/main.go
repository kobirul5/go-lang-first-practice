package main


type animal interface {
	Speak() string
}


type Dog struct {}
type Cat struct {}

func (d Dog) Speak() string {
	return "Woof!"
}

func (c Cat) Speak() string {
	return "Meow!"
}


func main() {
	var myDog Dog
	var myCat Cat

	println(myDog.Speak())
	println(myCat.Speak())

	// Output:
	// Woof!
	// Meow!
}
