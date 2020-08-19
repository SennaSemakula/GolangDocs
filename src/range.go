package main

import "fmt"

var primes = []int{2, 3, 5, 7, 9, 11, 13}
var evens = []int{0, 2, 4, 6, 8, 10}

func main() {
	for i, v := range primes {
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}

	// Grr I only want to iterate over the values. Easy!
	for _, v := range evens {
		fmt.Printf("Even number: %v\n", v)
	}
}
