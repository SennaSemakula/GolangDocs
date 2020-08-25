package main

import (
	"fmt"
	"time"
)

func count(s string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%v from: %v\n", s, i)
	}
}

func main() {
	count("sync")

	go count("goroutine")

	// Anonymous function
	go func(msg string) {
		fmt.Println(msg)
	}("going")
	
	time.Sleep(time.Second)
	fmt.Println("done")
}
