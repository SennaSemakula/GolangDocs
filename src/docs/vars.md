#Variables
var statement
- declares a list of variables
- can be defined at package or function level
- the type of the variable is declared at the end
- Can set an initialiser. If you omit the type, the variable will take the type of the initialiser
```
var languages python, javascript, golang, java, c = true, true, true, true, false
```

short variable declarations (:=)
- inside a function := short assignment can be used with implicit type
- Outside a function the := constrcut is not available
```
car := "Tesla"
```

Zero values
- variables declared without an explicit initial value are given their zero value
```
bool -> false
int -> 0
string -> "" (empty string)
```

constants
- declared using the const keyword
- cannot be declared using the := short assignment construct
```
const birthday = "04.05.2004"
