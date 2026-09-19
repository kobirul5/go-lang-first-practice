package main

import "fmt"

func makeCoffee(coffeeNo int) {
	fmt.Println("making Coffee....", coffeeNo)
}

func main() {

	// for i := 0; i <=10; i++{
	// 	makeCoffee(i)
	// }

	// // while styled loop
	// j := 0
	// for j <= 10 {
	// 	makeCoffee(j)
	// 	j++
	// }

	// break and continue 
	for i := 0; i <= 10; i++ {
		if i == 5 {
			continue
		}
		// if i == 8 {
		// 	break
		// }
		makeCoffee(i)
	}

}