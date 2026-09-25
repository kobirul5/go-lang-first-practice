package main

import (
	"fmt"
	"time"
)

func main() {

	var startTime = time.Now()

	// concurrency
	// uploadFile()
	// saveToDb()
	// sendEmail()

	// concurrency
	go uploadFile()
	go saveToDb()
	go sendEmail()

	time.Sleep(4* time.Second)


	fmt.Println("all task completed ")
	fmt.Println("Total time taken: ", time.Since(startTime))
}

func uploadFile() {
	// code to upload file
	fmt.Println("uploading file")
	time.Sleep(3 * time.Second)
	fmt.Println("file upload done ")
}

func saveToDb() {
	// code to save to database
	fmt.Println("saving to database")
	time.Sleep(3 * time.Second)
	fmt.Println("data saved to database ")
}
func sendEmail() {
	// code to send email
	fmt.Println("sending email")
	time.Sleep(3 * time.Second)
	fmt.Println("email sent ")
}
