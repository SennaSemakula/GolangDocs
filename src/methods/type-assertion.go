package main

import "fmt"

func GetType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("%v is an integer\n", v)
	case string:
		fmt.Printf("%v is a string\n", v)
	case bool:
		fmt.Printf("%v is a bool\n", v)
	default:
		fmt.Println("%v do not know the type", v)
	}
}

func main() {
	var i interface{} = "hello"

	t, ok := i.(string)

	fmt.Println(t, ok)

	// s := i.(float64) // panic will occur as we haven't specified bool value 
	// fmt.Println(s)

	GetType(21)
	GetType("apples")
	GetType(false)


}
