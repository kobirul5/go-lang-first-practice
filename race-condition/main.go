package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

var counter int


func main(){

for range 10000 {
	wg.Go(increment)
}

wg.Wait()

fmt.Println("Counter value is", counter)

}


func increment() {


	mu.Lock()
	counter = counter+1
	mu.Unlock()

}

// without mu.lock and mu. unlock result not perfect, thats why we use mu.lock and mu unlock