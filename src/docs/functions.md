# functions

- When two or more function parameters share a type, you can omit the type from all but the last:
```
func add(x int, y int) -> func add(x, y int)
```

- A function can return any number of results
```
func dob (day, month, year int) (int, int, int) {
    ...
}
```

Named return values
- Return values that are named

Naked return statements
- Should only be used in short functions. Can harm readability in longer functions

**Pass around values**
- Functions are values as well so they can be passed around 
- Functions may also be used as function arguments and return values
```
func compute(fn (x, y float64) float64) float64 {\
    ...
}
```

**Function closures**
- Go has closures which is a function that references variables from outisde is body
- The function may access and assign the referenced variables