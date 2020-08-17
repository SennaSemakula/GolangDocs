package main

import "fmt"

var python, javascript, golang, c = true, true, true, false

var mangoes, lemons, grapes, cucumber = "fruit", "fruit", "fruit", "vegetable"

const Pi = 3.14
const Truth = true

func main() {
	fmt.Println("Languages that I know:", python, javascript, golang, c)
	fmt.Println("Fruits:", mangoes, lemons, grapes, cucumber)

	car := "Tesla"
	fmt.Printf("I own a %v\n", car)

	fmt.Printf("Value of Pi is %v\n", Pi)
}