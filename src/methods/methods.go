package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type MyFloat float64

func (f MyFloat) AbsFloat() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}


func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	p := &v
	fmt.Println(p.Abs())

	// You can declare a method on non struct types as well
	f := MyFloat(-2)
	fmt.Println(f.AbsFloat())
}
