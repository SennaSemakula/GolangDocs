package main

import "fmt"

func swap(x, y *int) (*int, *int) {
	x = *y
	y = *x

	return x, y
}
func main() {
	*x := 1
	*y := 2
	fmt.Println("Swapped values are: ")
	fmt.Println(swap(&x, &y))
}
