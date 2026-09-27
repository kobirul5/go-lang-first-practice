package main

import (
	"fmt"
	"sync"
	"time"
)

//waitgroup is used to wait for a collection of goroutines to finish executing


var wg sync.WaitGroup


func main() {

	var startTime = time.Now()

	// concurrency
	// uploadFile()
	// saveToDb()
	// sendEmail()

	// concurrency
	wg.Add(1)
	go uploadFile()
	wg.Add(1)
	go saveToDb()
	wg.Add(1)
	go sendEmail()

	// time.Sleep(4* time.Second)
	wg.Wait() // wait for all the goroutines to finish executing


	fmt.Println("all task completed ")
	fmt.Println("Total time taken: ", time.Since(startTime))
}

func uploadFile() {
	// code to upload file
	fmt.Println("uploading file")
	time.Sleep(3 * time.Second)
	fmt.Println("file upload done ")
	wg.Done() // mark the goroutine as done
}

func saveToDb() {
	// code to save to database
	fmt.Println("saving to database")
	time.Sleep(3 * time.Second)
	fmt.Println("data saved to database ")
	wg.Done() // mark the goroutine as done
}
func sendEmail() {
	// code to send email
	fmt.Println("sending email")
	time.Sleep(3 * time.Second)
	fmt.Println("email sent ")
	wg.Done() // mark the goroutine as done
}


// Output:
// uploading file
// saving to database
// sending email
// file upload done
// data saved to database
// email sent
// all task completed
// Total time taken: 3.000123456s	

