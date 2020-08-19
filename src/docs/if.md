#If statements
- Similar to for loops where there are no parentheses. {} braces are required
```
// 
if x%2 == 0 {
    // even number
}
```
- Go has a short if statement which similiar to for it can star with a short statement before execution
- variables declared in a short if statement are also available in else blocks
- cannot use short if variable outside blocks
```
if v:= 2020; v < 2020 {
    return "You're in the past"
}
```