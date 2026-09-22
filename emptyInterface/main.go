package main

import "fmt"

func Process(data any){
  staDate, ok := data.(string)
  if ok {
	fmt.Println("data is a string", staDate)
  }

  intDate, ok := data.(int)
  if ok {
	fmt.Println("data is not an int", intDate + 100)
  }


}


func main() {}