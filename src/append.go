package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s)

	s = append(s, 1)
	fmt.Println(s)
	printSlice(s)
	s = append(s, 2, 3, 5, 7, 9) // primes
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("Length: %d, Capacity: %d Items: %v\n", len(s), cap(s), s)
}
