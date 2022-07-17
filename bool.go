package main

import (
	"fmt"
)

func replaceArrPart[E any](arr []E, start, end int, repl E) []E {
	var newArr []E
	for i := 0; i < start; i++ {
		newArr = append(newArr, arr[i])
	}
	newArr = append(newArr, repl)
	for i := end + 1; i < len(arr); i++ {
		newArr = append(newArr, arr[i])
	}
	return newArr
}

func containsArr[E comparable](arr []E, obj E) bool {
	for _, v := range arr {
		if v == obj {
			return true
		}
	}
	return false
}

func BoolEval(str string, tested bool) bool {
	// ~(true & true) | false
	fmt.Println("before test", str)

	if tested == false {
		str = Space.ReplaceAllString(str, "")
		nest := nestedPar(str)
		maxdp := max(nest)
		if maxdp > 0 {
			start, end := getPart(nest, maxdp)
			str = replacePart(str, fmt.Sprint(BoolEval(str[start:end], tested)), start-1, end+1)
			return BoolEval(str, tested)
		} else {
			tested = true
		}
	}

	// ~(true & true | false) | false
	// true & true | false
	alo := Lex(str)
	fmt.Println("protoalo", alo)
	for containsArr(alo, tok{"AND", "&"}) || containsArr(alo, tok{"OR", "|"}) || containsArr(alo, tok{"NOT", "~"}) {
		for i := 0; i < len(alo); i++ {
			if alo[i].typ == "AND" {
				if alo[i-1].value == "true" && alo[i+1].value == "true" {
					alo = replaceArrPart(alo, i-1, i+1, tok{"KEYWORD", "true"})
				} else {
					alo = replaceArrPart(alo, i-1, i+1, tok{"KEYWORD", "false"})
				}
			} else if alo[i].typ == "OR" {
				if alo[i-1].value == "true" || alo[i+1].value == "true" {
					alo = replaceArrPart(alo, i-1, i+1, tok{"KEYWORD", "true"})
				} else {
					alo = replaceArrPart(alo, i-1, i+1, tok{"KEYWORD", "false"})
				}
			} else if alo[i].typ == "NOT" {
				if alo[i+1].value == "true" {
					alo = replaceArrPart(alo, i, i+1, tok{"KEYWORD", "false"})
				} else if alo[i+1].value == "false" {
					alo = replaceArrPart(alo, i, i+1, tok{"KEYWORD", "true"})
				}
			}
		}
		fmt.Println("alo", alo)
	}
	if alo[0].value == "true" {
		return true
	}

	return false
}
