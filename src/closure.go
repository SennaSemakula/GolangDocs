package main

import "fmt"

func adder() func(int) int {
	sum := 0

	return func(x int) int {
		sum += 1
		return sum
	}
}
func main() {
	incr := adder()

	for i := 0; i < 10; i++ {
		fmt.Println(incr(i))
	}
}
