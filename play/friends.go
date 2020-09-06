package main

import "fmt"

var friends = make(map[string]string)


func newFriends(f *[]string) {
	vals := *f
	fmt.Println("Moved to university so making new friends!")
	vals[0] = "Jim"
	vals[1] = "Katey"
}
func main() {
	// creates a slice
	myFriends := []string{"Senna", "Blake"}
	// myFriendsPtr := &myFriends
	for _, name := range myFriends {
		fmt.Printf("%v is a friend\n", name)
	}
	newFriends(&myFriends)

	for _, name := range myFriends {
		fmt.Printf("%v is a friend\n", name)
	}
}
