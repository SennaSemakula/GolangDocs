# Methods
- methods are functions with a special receiver argument
- you can call methods on different types e.g. structs, floats
```
func (s Dog) Bark() int {
    return "WOOF!"
}

dog := Dog{"Bill", "Brown", 2}
fmt.Println("Dog goes...", dog.Bark())
```
- remember methods are just functions with a receiver argument

**pointer receivers**
- Methods with pointer receivers can modify the value of the receiver
- As methods often need to modify the value of the reciver, pointer receivers are more common that value receivers
- If you use value receivers, the method will only copy the original value of the reciever; not modify it
```
func (v *T) Receiver()
```

**POINTER INDIRECTION**
- Functions with a pointer argument must take a pointer or it will throw a compile error
- methods with pointer receivers can either take a value or a pointer
```
func (v *Vertex) Scale(f float64) {}

var v Vertex
v.Scale(5) // OK >> Go automatically interprets it as v.scale(&5)

p := &v
p.Scale(10) // OK
```
- Works the other way round with value arguments and receivers
- Methods should either have value or pointer receivers, not both!

Questions?
- Why use pointer receivers over value ones?
    * so that the method can modify the value the receiver points to
    * to avoid copying the value on each method call. this can be more efficient if the receiver is a large struct

