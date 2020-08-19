- Go's switch statement only runs the selected case not all cases that follow unlike Java, C, C++, PHP and JavaScript
- Go also doesn't require a break statement at the end of each case as it does it automatically unlike the other languages
- switch cases evaluate from top to bottom e.g. if the first case passes, the otther cases will not run
- a switch with no condition will also evaluate to true. Good way to write long if-then-else chains

```
switch fruit := "Lemon", fruit {
    case "Lemon":
    fmt.Println("I like.")
    case "Apple"
    fmt.Println("Do not like.")
    default:
    fmt.Println("I love %v\n", fruit)
}
```