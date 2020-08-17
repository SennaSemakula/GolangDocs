package main

import (
	"fmt"
	"math"
)

func add(x, y int) int {
	return x + y
}

func substract(x, y int) int {
	return x - y
}

func divide(x, y int) int {
	return x / y
}

func dob(day, month, year int) (int, int, int) {
	return day, month, year
}

// naked return function
func split(sum int) (x, y int) {
	x = sum / 2
	y = sum - x

	return
}

func main() {
	fmt.Println("Hello World from Senna")
	fmt.Printf("Square root of 9 is %g\n", math.Sqrt(9))

	day, month, year := dob(5, 05, 2000)
	fmt.Println("Date of birth:", day, "/", month, "/", year)
	fmt.Println(split(20))
}