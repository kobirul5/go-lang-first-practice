package main
import "fmt"

func main(){

	day := "Monday"

	switch day { // target variable
	case "Monday":
		fmt.Println("Today is Monday.")
	case "Tuesday":
		fmt.Println("Today is Tuesday.")
	case "Wednesday":
		fmt.Println("Today is Wednesday.")
	case "Thursday":
		fmt.Println("Today is Thursday.")
	case "Friday":
		fmt.Println("Today is Friday.")
	default:
		fmt.Println("It's the weekend!")
	}

	// normal switch statement
	age := 70
	switch {
	case age < 18:
		fmt.Println("You are a minor.")
	case age >= 18 && age < 65:
		fmt.Println("You are an adult.")
	default:
		fmt.Println("You are a senior citizen.")
	}


}