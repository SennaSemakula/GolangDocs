#Arrays
- You can create a fixed size array by using the following
```
var a [10]int // initialises an array with 0s
a := [3]int{0, 0, 0}// short hand
```
**slices**
- You can select values from arrays by [low: high] slice
- It will include the first element by exclude the last one
- Slices are dynamically sized
- **slices are like references to arrays. changing elements of a slice, modifies the contents of the underlying array**
- You can create dynamically sized arrays by using slice literals
```
students := []struct{
    name string
    age int
    gender string
    enrolled bool
}{
    {"James", 22, "Male", true},
    {"Fiona", 23, "Female", true},
}
```

**len and capacity of slices**
- You can call s items len(s) to get the total items of the slice and cap(s) to see the number of items the underlying array has
- you can extend the slices capacity
- NOTE: if you extend a slice beyond it's capacity, a panic with a bounds out of range error will be thrown
- A slice with length and capacity 0 will return nil

**create dynamic sized  slice with make function**
- make()

**append function**
- you can append items to a slice using the built in package
- if the slice has sufficient capacity it wille xtend the slice, if not, it will use the underlying array
- append returns a new slice that is why it's neccessary to store the result of an append
```
append(slice []type, elems...Type)
```

**range function**
- The range function iterates over a slice or map in a for loop
- when ranging over a slice, two items are returns; indeex and value
```
for i, v := range elems {
    fmt.Printf("Index: %v, Value: %v", i, v)
}
```

