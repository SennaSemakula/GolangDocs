package main

import "fmt"

func slice(start, end int, arr [10]int) []int {
	v := arr[start:end]
	return v
}

func isEmpty(arr []string) bool {
	if arr == nil {
		fmt.Println("array is empty")
		return false
	}
	fmt.Println("Array is not empty")
	fmt.Println("Capacity", len(arr))
	return true
}

func main() {
	var numbers [10]int

	//short hand assignment of arrays
	fruits := [3]string{"grapes", "passion fruit", "oranges"}

	// Slicing an array will create dynamically sized
	var myFruits []string = fruits[0:2]
	numbers[2] = 22
	numbers[3] = 10
	fmt.Println(numbers)
	fmt.Println(slice(0, 1, numbers))
	fmt.Println("My favourite fruits:", myFruits)

	//slice literals
	b := []bool{true, true, false}
	fmt.Println("Slice literal:", b)

	//Create a slice of structs
	students := []struct {
		name     string
		age      int
		gender   string
		enrolled bool
	}{
		{"Blake", 22, "Male", true},
		{"Fiona", 21, "Female", true},
		{"Griffin", 18, "Male", false},
	}
	fmt.Println("Students in database:", students)

	var emptySlice []string
	isEmpty(emptySlice)

	a := make([]int, 5) // will create dynamically sized array array. 3rd argument allows you specify the capacity
	bravo := make([]int, 0, 5)
	c := bravo[:2]
	fmt.Println(a)
	fmt.Println(len(bravo))
	fmt.Println(c)
}
