package main

import (
	"fmt"
	"io"
)

func Reverse(r io.Reader) io.Reader {
	b := make([]byte, 8)

	for {
		n, err = r.Read(b)
		fmt.Println(n)

		if err == io.EOF {
			break
		}
	}
	return b
}
func main() {
	r := strings.NewReader("Go is a multi purpose programming language")

}