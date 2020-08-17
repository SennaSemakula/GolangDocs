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