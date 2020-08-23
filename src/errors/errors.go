package main

import (
	"fmt"
	"time"
)

// We want to override the Error() method
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("At %v. Cause: %v", e.When, e.What)
}

func run() error{
	return &MyError{time.Now(), "Error occured!"}
}


func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}