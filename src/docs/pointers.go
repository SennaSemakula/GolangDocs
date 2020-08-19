# Pointers
Go has pointers that holds the memory address of a value
```
i, j := 100, 200
p := &i // set pointer to i memory address
*p = 200 // change value of pointer (i will now be 200)
fmt.Println(i) // should print 200
```
- you can derence a pointer to access its value using the * operator
- to get the memory address of a pointer you can use the & operator
- Unlike C, Go has no pointer arithmetic 