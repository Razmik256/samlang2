# samlang2

```coffeescript
x = [1,2,3]; #declare variable
x = x ^ 1; #add to array
x = x!!0; #get array index value
x = "alo";
x = x!!0; #strings are also arrays

#comparison
if true {
  x = 0;
}

#while
i = 0;
while i < 10 {
  i = i + 1;
}

#pure function declaration
func alo(x,y) {
  return x + y;
}
z = alo(2,3);
```

```coffeescript
x = true;
x = 2;
y = "alo";
z = [1,2,3];

while x < 10 {
  x = x + 1;
  if x > 5 {
    break;
  }
}

func alo(x,y) {
  return x + y;
}
```
Fibonnaci
```coffeescript
func fib(max) {
    arr = [];
    num1 = 0;
    num2 = 1;
    i = 0;
    while i < max {
        a = num2;
        num2 = num1 + num2;
        num1 = num2;
        arr = arr ^ num2;
        i = i + 1;
    }
    return arr;
}
z = fib(10);
```

- [x] numbers
- [x] booleans
- [x] arrays
- [x] strings
- [x] if else
- [x] loops
- [x] functions
- [ ] subroutines
- [x] local variables
- [ ] memory managment
