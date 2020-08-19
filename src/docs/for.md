# For
- Go has only one looping construct: For
- Unlike languages like C, Java and JavaScript, the for loop in go has no parentheses
- The init and post statements are optional e.g.
- drop the semicolons for a while loop
- omitting the loop condition will create an infinite loop
```
for i := 0; i < 10; i++ {

}

```
```
sum := 1
for ; sum < 10; {

}

```

**for is Go's while**
```
sum := 0
for sum < 1000 {
    // will cause an infinite loop
}
