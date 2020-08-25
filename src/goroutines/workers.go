package main

import (
	"fmt"
	"time"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	fmt.Printf("%d worker starting now\n", id)
	time.Sleep(time.Second)
	fmt.Printf("%d worker finishing\n", id)

	wg.Done() // decrement the counter by one
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1) // add each worker to the waitgroup. Must be done beforee calling the goroutine
		go worker(i, &wg)
	}

	// block goroutines until all have finished
	wg.Wait()

}