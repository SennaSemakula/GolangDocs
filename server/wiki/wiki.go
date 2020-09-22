package wiki

import (
	"fmt"
	"io/ioutil"
	"log"
)

type Page struct {
	Title string
	Body  []byte // slice of bytes
}

func LoadPage(title string) (*Page, error) {
	filename := title + ".txt"

	body, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

// Receiver called save that takes in a Page struct
func (p *Page) Save() error {

	// read/write permissions for the current user
	err := ioutil.WriteFile(p.Title+".txt", p.Body, 0600)

	if err != nil{
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
