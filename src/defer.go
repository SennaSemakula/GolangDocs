package main

import "fmt"

func main() {
	defer fmt.Println("Hello Senna")

	println("Good morning Senna")

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
