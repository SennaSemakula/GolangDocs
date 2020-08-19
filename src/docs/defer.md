# Defer
- defer statement defers the execution of a function until the surronding function returns
- the deferred call's arguments are evaluated immediately but the function call is not executed until the surrounding function returns
- deffered function calls are pushed onto a stack. when a function returns, its deferred calls are executed last-in-first-out-order
- read https://blog.golang.org/defer-panic-and-recover
```
defer fmt.Println("world") //arguments are evaluated
// this will be called first
fmt.Println("world")
```
- What will the following print out?
```
fmt.Println("counting")

for i := 0; i < 10; i++ {
    defer fmt.Println(i)
}

fmt.Println("done")
```