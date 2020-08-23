#Errors
- Go programs express error state with error values
- The error type is a built in interface similar to ```fmt.Stringer```
- Functions often return an error value. You should test for errors to see if it's != nil
```
type error interface {
    Error() string
}

Example:
```
i, err := strconv.Atoi("56")
if err != nil {
    fmt.Printf("Cannot convert number %v to int\n., err)
    return
}
fmt.Println("Converted integer", i)
```
