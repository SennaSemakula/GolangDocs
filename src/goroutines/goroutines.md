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
