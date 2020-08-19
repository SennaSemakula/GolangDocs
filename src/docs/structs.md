#Structs
- collection of fields
- you can access structs with the . notion
- you can have pointers to structs e.g.
```
v := Struct{1, 2}
p := &v
p.X = 100
fmt.Println("Memory address of struct is, &p)
```