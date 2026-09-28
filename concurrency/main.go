package main

import (
	"fmt"
	"sync"
	"time"
)

//waitgroup is used to wait for a collection of goroutines to finish executing


var wg sync.WaitGroup


// func main() {

// 	var startTime = time.Now()

// 	// concurrency
// 	// uploadFile()
// 	// saveToDb()
// 	// sendEmail()

// 	// concurrency
// 	wg.Add(1)
// 	go uploadFile()
// 	wg.Add(1)
// 	go saveToDb()
// 	wg.Add(1)
// 	go sendEmail()

// 	// time.Sleep(4* time.Second)
// 	wg.Wait() // wait for all the goroutines to finish executing


// 	fmt.Println("all task completed ")
// 	fmt.Println("Total time taken: ", time.Since(startTime))
// }

func main() {

	var startTime = time.Now()

	wg.Go(uploadFile)
	wg.Go(saveToDb) 
	wg.Go(sendEmail)


	wg.Wait() 

	fmt.Println("all task completed ")
	fmt.Println("Total time taken: ", time.Since(startTime))
}

func uploadFile() {
	// code to upload file

	// defer wg.Done() 

	fmt.Println("uploading file")
	time.Sleep(3 * time.Second)
	fmt.Println("file upload done ")
	// wg.Add(-1)
	// wg.Done() // mark the goroutine as done
}

func saveToDb() {
	// code to save to database
	// defer wg.Done() 
	fmt.Println("saving to database")
	time.Sleep(3 * time.Second)
	fmt.Println("data saved to database ")
		// wg.Add(-1)
	// wg.Done() // mark the goroutine as done
}
func sendEmail() {
	// code to send email]
	// defer wg.Done()
	fmt.Println("sending email")
	time.Sleep(3 * time.Second)
	fmt.Println("email sent ")
		// wg.Add(-1)
	// wg.Done() // mark the goroutine as done when the function exits
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

