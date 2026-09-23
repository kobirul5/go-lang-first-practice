package main

import "fmt"

// Variadic function
func Sum(nums ...int) int{

	  total := 0
	  for _, n := range nums {
		  total += n
	  }
	  return total
}


func greet(name string, msg ...string) string {

	for _, m := range msg {
		fmt.Println(m, name)
	}
	return name


}


func main() {
	sum := Sum(1, 2, 3, 4, 5)
	println(sum)

	greet("welcome", "jamal", "kamal", "jhon deo")


}