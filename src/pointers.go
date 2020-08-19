package main

import "fmt"

func main() {
	i, j := 21, 100

	p := &i // point p to i
	*p = 1202 // set i through pointer p
	fmt.Println(*p)
	fmt.Println(j)
	fmt.Printf("Value of i is %v\n", *p)
	p = &j //point to j memory address
	*p = 200
	fmt.Printf("Updated value of j is %v\n", j)


}
