package main

import (
	"fmt"
	"time"
)



func main(){

	var ch = make(chan string)
	go UploadFile(ch)
	
}


func UploadFile(c chan string){
	fmt.Println("uploading file....")
	time.Sleep(3*time.Second)
	fmt.Println("File upload done!")
	c <- "https://s3.3443432.png"
}
