#Interfaces
- An interface type is defined as a set of method signatures
- It's a set of methods but also a type 
- There is no implements keyword in Go
- When passing an interface into a function's argument it will have only one value at runtime; interface
    - At runtime go does type conversion to ensure that only one value is at runtime
```
type Dog interface {
    Bark()
    Walk()
    Eat()
}
```
- interfaces have a value and a concrete tye (v, T) tuple
    - this allows the interface to determine what method to call based on v, T
- if the concrete value is nil, the method will be called with a nil receiver only if the type is not nil
    - **in other languages this will trigger a null pointer exception**
- if an interface that hold both value and type is nil, it'll trigger a runtime exception
    - this is because there is not type inside the interface tuple to indicate what method to call

**empty interface**
- implements zero methods

Examples
```
func PrintAll(vals []interface{}) {
    for _, val := range vals {
        fmt.Println(val)
    }
}

func main() {
    //wrong because expects a slice of interface type
    names := []string{"stanley", "david", "oscar"}

    //correct 
    //names := make([]interface{}, len(names))
    PrintAll(names)

    for i, v := range names {
        vals[i] = v
    }
    PrintAll(vals)
}
```


