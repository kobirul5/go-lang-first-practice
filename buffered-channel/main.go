package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan string, 3)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "file upload"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		ch <- "file url save"
	}()
	go func() {
		time.Sleep(3 * time.Second)
		ch <- "email sent"
	}()

	for range 3 {
		date := <-ch
		fmt.Println(date)
	}

}
