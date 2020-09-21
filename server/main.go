package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

type Page struct {
	title string
	body  []byte // slice of bytes
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"

	body, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return &Page{title: title, body: body }, nil
}

// Receiver caled save that takes in a Page struct
func (p *Page) save() error {

	// read/write permissions for the current user
	err := ioutil.WriteFile(p.title + ".txt", p.body, 0600)

	if err != nil {
		log.Fatal(err)
	}

	return err
}

func ReadFile(title string) ([]byte, error) {
	contents, err := ioutil.ReadFile(title)

	fmt.Printf("Attempting to read file %v\n", title)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File contents: %v\n", contents)
	return contents, err
}
func main() {
	p1 := &Page{title: "Greetings", body: []byte("Hello world to all my Gophers!")}

	p1.save()

	p2, _ := loadPage("Greetings")
	fmt.Println(string(p2.body))

}
