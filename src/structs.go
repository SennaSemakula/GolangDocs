package main

import "fmt"

type Car struct {
	Make  string
	Model string
	Age   int
	Color string
}

func main() {
	car := Car{"Tesla", "X", 2020, "Black"}
	fmt.Printf("The make of my car is %v\n", car.Make)
	// Pointer to struct
	p := &car
	p.Color = "Blue"
	fmt.Println("Pointer to car colour is", car.Color)
	fmt.Println("Memory addr of struct is ", &p)
}
