package main

import (
	"fmt"
	"strconv"
)

func parseFloatNoErr(str string) float64 {
	a, err := strconv.ParseFloat(str, 64)
	if err != nil {
		panic(err)
	}
	return a
}

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

func BoolEval(str string, tested bool) string {
	// ~(true & true) | false

	// true false--------------------------------
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
	lex := Lex(str)
	fmt.Println(lex)
	// arrays
	for i := 0; i < len(lex); i++ {
		if lex[i].value == "[" {
			start := i
			a := "["
			for lex[i].typ != "GREATER" {
				i++
				a += lex[i].value
			}
			v := ArrEval(a)
			lex = replaceArrPart(lex, start, i, tok{"KEYWORD", v[1 : len(v)-1]})
		} else if lex[i].value == "(" {
			// function
			fmt.Println(lex)
		} else if (lex[i].value == "+" && lex[i+1].value != "+") || lex[i].value == "-" || lex[i].value == "/" || lex[i].value == "*" {
			lex = replaceArrPart(lex, i-1, i+1, tok{"KEYWORD", fmt.Sprintf("%g", MathEval(lex[i-1].value+lex[i].value+lex[i+1].value, false))})
		}
	}
	for i := 1; i < len(lex)-1; i++ {
		if lex[i].typ == "LESS" {
			if parseFloatNoErr(lex[i-1].value) < parseFloatNoErr(lex[i+1].value) {
				lex = replaceArrPart(lex, i-1, i+1, tok{"KEYWORD", "true"})
			} else {
				lex = replaceArrPart(lex, i-1, i+1, tok{"KEYWORD", "false"})
			}
		} else if lex[i].typ == "GREATER" {
			if parseFloatNoErr(lex[i-1].value) > parseFloatNoErr(lex[i+1].value) {
				lex = replaceArrPart(lex, i-1, i+1, tok{"KEYWORD", "true"})
			} else {
				lex = replaceArrPart(lex, i-1, i+1, tok{"KEYWORD", "false"})
			}
		} else if lex[i].typ == "EQUAL" && lex[i+1].typ == "EQUAL" {
			if parseFloatNoErr(lex[i-1].value) == parseFloatNoErr(lex[i+2].value) {
				lex = replaceArrPart(lex, i-1, i+2, tok{"KEYWORD", "true"})
			} else {
				lex = replaceArrPart(lex, i-1, i+2, tok{"KEYWORD", "false"})
			}
		} else if lex[i].typ == "NOT" && lex[i+1].typ == "EQUAL" {
			if parseFloatNoErr(lex[i-1].value) != parseFloatNoErr(lex[i+2].value) {
				lex = replaceArrPart(lex, i-1, i+2, tok{"KEYWORD", "true"})
			} else {
				lex = replaceArrPart(lex, i-1, i+2, tok{"KEYWORD", "false"})
			}
		}
	}
	for i := 0; i < len(lex); i++ {
		if lex[i].typ == "NOT" {
			if lex[i+1].value == "true" {
				lex = replaceArrPart(lex, i, i+1, tok{"KEYWORD", "false"})
			} else if lex[i+1].value == "false" {
				lex = replaceArrPart(lex, i, i+1, tok{"KEYWORD", "true"})
			}
		}
	}
	for i := 0; i < len(lex); i++ {
		if lex[i].typ == "AND" {
			if lex[i-1].value == "true" && lex[i+1].value == "true" {
				lex = replaceArrPart(lex, i-1, i+1, tok{"KEYWORD", "true"})
			} else {
				lex = replaceArrPart(lex, i-1, i+1, tok{"KEYWORD", "false"})
			}
		}
	}
	for i := 0; i < len(lex); i++ {
		if lex[i].typ == "OR" {
			if lex[i-1].value == "true" || lex[i+1].value == "true" {
				lex = replaceArrPart(lex, i-1, i+1, tok{"KEYWORD", "true"})
			} else {
				lex = replaceArrPart(lex, i-1, i+1, tok{"KEYWORD", "false"})
			}
		}
	}
	if lex[0].value == "false" {
		return "false"
	} else if lex[0].value == "true" {
		return "true"
	} else {
		return lex[0].value
	}
}
