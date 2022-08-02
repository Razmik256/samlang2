# samlang2

```ruby
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
```
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
