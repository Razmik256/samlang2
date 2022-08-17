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

```coffeescript

func changearr(arr, arrlen, index, value) {
    newarr = [];
    arrstart = [];
    newarr = [];
    i = 0;
    if index > 0 {
        while i < index {
            arrstart = arrstart ^ (arr!!i);
            i = i + 1;
        }
        newarr = (newarr ^ arrstart) ^ value;
        i = index+1;
    } else {
        i = i + 1;
        newarr = newarr ^ value;
    }
    arrend = [];
    while i < arrlen {
        arrend = arrend ^ (arr!!i);
        i = i + 1;
    }
    newarr = newarr ^ arrend;
    return newarr;
}
func swap(arr, arrlen, index1, index2) {
    ind = arr!!index1;
    arr = changearr(arr, arrlen, index1, arr!!index2);
    arr = changearr(arr, arrlen, index2, ind);
    return arr;
}
func bubblesort(arr, len) {
    i = 0;
    j = 0;
    while i < (len-1) {
        while j < (len-i-1) {
            if arr!!j > arr!!(j+1) {
                arr = swap(arr, len , j, j+1);
            }
            j = j + 1;
        }
        j = 0;
        i = i + 1;
    }
    return arr;
}
a = bubblesort([5577006791947779410,8674665223082153551,6129484611666145821,4037200794235010051,3916589616287113937,6334824724549167320,605394647632969758,1443635317331776148,894385949183117216,2775422040480279449,4751997750760398084,7504504064263669287,1976235410884491574,3510942875414458836,2933568871211445515,4324745483838182873,2610529275472644968,2703387474910584091,6263450610539110790,2015796113853353331,1874068156324778273,3328451335138149956,5263531936693774911,7955079406183515637,2703501726821866378,2740103009342231109,6941261091797652072,1905388747193831650,7981306761429961588,6426100070888298971,4831389563158288344,261049867304784443], 32);
```

- [x] numbers
- [x] booleans
- [x] arrays
- [x] strings
- [x] if else
- [x] loops
- [x] functions
- [ ] solve broken return statement
- [ ] subroutines
- [x] local variables
- [ ] memory managment
