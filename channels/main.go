package main

import (
	"fmt"
	"time"
)



func main(){

	var ch = make(chan string)
	go UploadFile(ch)
	fileUrl := <-ch

	fmt.Println(fileUrl)
}


func UploadFile(c chan string){
	fmt.Println("uploading file....")
	time.Sleep(3*time.Second)
	fmt.Println("File upload done!")
	c <- "https://s3.3443432.png"
}
