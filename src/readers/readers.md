# Readers
- Readers is any type that implements the 'Read()' method
- An example could be when you open a file with ```os.Open()```. The object returned is ```os.File``` which is a Reader because it implments the Read method.

**How it works**
- You pass in a slice of bytes which the Reader is then asked to populate it with data until it runs out of data
    - It returns the number of bytes read (n) or an error
    - After it has finished reading, it will return a special marker ```io.EOF```

**Examples**
```
r := strings.NewReader("Hello, Reader!") // Returns a reader object
b := make([]byte, 8)

for {
    // will keep incrementing and reading the bytes of r until is satifies the length of b
    n, err := r.Read(b)

    if err == io.EOF { // finished reading
        break
    }
}

r.Read()
```
- returns a io.Reader object which you can 

NOTE: https://medium.com/@matryer/golang-advent-calendar-day-seventeen-io-reader-in-depth-6f744bb4320b