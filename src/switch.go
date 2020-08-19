package main

import (
	"fmt"
	"runtime"
	"time"
)

func getOS() {
	fmt.Println("Go is currently running on\n")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OSX")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s", os)
	}
}

func whensSat() time.Weekday {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("It's today!")
	case today + 1:
		fmt.Println("It's tommorrow!")
	case today + 2:
		fmt.Println("It's in two days!")
	default:
		fmt.Println("It's too far away!")
	}

	return today

}

func main() {
	getOS()
	whensSat()

	// Better way of writing long if-then-else chains
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
