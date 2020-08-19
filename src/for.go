package main

import "fmt"

func loop1() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func loop2() int {
	sum := 1
	for sum < 10 {
		sum += sum
	}

	return sum
}

func whileloop() int {
	sum := 0
	for sum < 1000 {
		fmt.Println(sum)
	}

	return sum
}

func main() {
	sum := 1
	for sum < 10 {
		sum += sum
	}
	fmt.Printf("New loop is %v\n", loop2())
	fmt.Println(sum)
	whileloop()
}
