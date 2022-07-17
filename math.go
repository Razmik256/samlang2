package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func nestedPar(str string) []int {
	str = strings.TrimSpace(str)
	var depth int = 0
	arr := []int{}
	for i := 0; i < len(str); i++ {
		if string(str[i]) == "(" {
			depth += 1
		} else if string(str[i]) == ")" {
			depth -= 1
		}
		arr = append(arr, depth)
	}
	return arr
}

func max(arr []int) int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

func getPart(arr []int, num int) (int, int) {
	start := -1
	end := len(arr)
	for i := 0; i < len(arr); i++ {
		if arr[i] == num {
			start = i + 1
			break
		}
	}
	if start != -1 {
		for i := start; i < len(arr); i++ {
			if arr[i] != num {
				end = i
				break
			}
		}
	}
	return start, end
}

func replacePart(str, repl string, start, end int) string {
	return strings.Replace(str, str[start:end], repl, 1)
}

var Space = regexp.MustCompile(`\s+`)

func MathEval(str string, t bool) float64 {
	tested := t
	var ret float64 = 0
	// if it's just a number return it
	if !strings.Contains(str, "+") &&
		!strings.Contains(str, "*") &&
		!strings.Contains(str, "-") &&
		!strings.Contains(str, "/") &&
		!strings.Contains(str, "(") &&
		!strings.Contains(str, ")") &&
		!strings.Contains(str, " ") &&
		str != "" {
		a, err := strconv.ParseFloat(str, 64)
		if err != nil {
			fmt.Println("wrong math value " + str)
			panic(err)
		}
		return a

	}

	// evaluate by parenth depth

	// [0,0,0,0,0,1,1,1,1,1,2,2,2,2,2,1,1,1,1]
	if tested == false {
		str = Space.ReplaceAllString(str, "")
		nest := nestedPar(str)
		maxdp := max(nest)
		if maxdp > 0 {
			start, end := getPart(nest, maxdp)
			str = replacePart(str, fmt.Sprint(MathEval(str[start:end], tested)), start-1, end+1)
			return MathEval(str, tested)
		} else {
			tested = true
		}
	}

	// iterate trough signs and make an evaluation tree
	for i := 0; i < len(str); i++ {
		if string(str[i]) == "+" {
			arr := strings.Split(str, "+")
			for i := 0; i < len(arr); i++ {
				ret += MathEval(arr[i], tested)
			}
			// fmt.Println("ret: " + fmt.Sprint(ret))
			return ret
		}
	}
	for i := 0; i < len(str); i++ {
		if string(str[i]) == "-" {
			arr := strings.Split(str, "-")
			ret = MathEval(arr[0], tested)
			for i := 1; i < len(arr); i++ {
				ret -= MathEval(arr[i], tested)
			}
			fmt.Println("mathret: " + fmt.Sprint(ret))
			return ret
		}
	}
	for i := 0; i < len(str); i++ {
		if string(str[i]) == "*" {
			arr := strings.Split(str, "*")
			ret = 1
			for i := 0; i < len(arr); i++ {
				ret *= MathEval(arr[i], tested)
			}
			// fmt.Println("ret: " + fmt.Sprint(ret))
			return ret
		}
	}
	for i := 0; i < len(str); i++ {
		if string(str[i]) == "/" {
			arr := strings.Split(str, "/")
			ret = MathEval(arr[0], tested)
			for i := 1; i < len(arr); i++ {
				ret /= MathEval(arr[i], tested)
			}
			// fmt.Println("ret: " + fmt.Sprint(ret))
			return ret
		}
	}
	return ret
}
