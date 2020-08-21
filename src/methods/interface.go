package main

import "fmt"

type Animal interface {
	Run()
	Speak() string
}

type Bird struct {
}

type Tiger struct {
}

func (t Tiger) Run() {
	fmt.Println("Running now!")
}

func (b Bird) Run() {
	fmt.Println("I can't run. Only fly!!!!")
}

func (t Tiger) Speak() string {
	return "ROAAAAARR!!!"
}

func (b Bird) Speak() string {
	return "Peck peck peck peck"
}

func Hunt(a Animal) {
	a.Run()
	fmt.Println(a.Speak())
}

func main() {
	animals := []Animal{Tiger{}, Bird{}}

	for _, animal := range animals {
		fmt.Println(animal.Speak())
		animal.Run()
	}

	// Hunt(parrot)
	// Hunt(tiger)
}
