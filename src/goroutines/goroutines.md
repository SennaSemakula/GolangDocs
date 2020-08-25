# Goroutines
- A goroutine is a lightweight thread of execution
- It allows us to run funciton calls asynchronously 
- You can call a function in a goroutine eby prefixing it with the go keyword beforehand
**NOTE:**
without goroutines we will have blocking function calls which will slow down the execution of our program

# Anonymous functions
- Functions what can be called without a function name
```
go func (msg string) {
    fmt.Println(msg)
}("hello world")
```

# Wait groups
- WaitGroup type from sync package is used gto wait for a collection of Goroutines. Can also be used to block all the Goroutines until all have finished
- The main goroutine calls the add method to set the number of Goroutines to wait for
- Call Add() to add a Goroutine into the WaitGroup.
    - typically this should be called before the statement that creates the goroutine because calls may happen at any time if the counter is greater than 0 or negative. 
- Once the counter of the WaitGroup goes back down to zero, all the Goroutines blocked in the WaitGroup are unblocked
**Example**:
```
package main

import (
	"sync"
)

type httpPkg struct{}

func (httpPkg) Get(url string) {}

var http httpPkg

func main() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}
	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL.
			http.Get(url)
		}(url)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
}
```
