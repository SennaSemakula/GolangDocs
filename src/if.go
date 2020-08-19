package main

import (
	"fmt"
	"math"
)

func isEven(x int) bool {
	if x%2 == 0 {
		return true
	} else {
		return false
	}
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func shortIf(year int) string {
	result := ""
	if v := year; v < 2020 {
		result = "Past"
	} else {
		result = "Future"
	}
	return result
}

func main() {
	fmt.Println(isEven(4))
	fmt.Printf("I think you're in the %v\n", shortIf(2030))
}
