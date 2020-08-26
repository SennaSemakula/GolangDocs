package main

import "fmt"

func greet(c chan string) {
	names := []string{
		"Senna",
		"Perry",
		"Grant",
	}

	for _, name := range names {
		c <- name
	}
	close(c) // close the channel
}

func main() {
	fmt.Println("Main Goroutine has started")
	c := make(chan string) // create the channel

	go greet(c)

	for {
		val, ok := <-c 
		if ok == false {
			fmt.Println(val, ok, "No more values to read. Breaking out of loop")
			break
		} else {
			fmt.Println("Greeting from", val)
		}
	}

	fmt.Println("Main Goroutine has finished")
}