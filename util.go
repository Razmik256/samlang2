package main

import (
	"fmt"
	"strconv"
	"strings"
)

func interfaceIntoString(v variable) string {
	switch v.typ {
	case "NUMBER":
		a, err := strconv.ParseFloat(v.value.(string), 64)
		if err != nil {
			panic(err)
		}
		return fmt.Sprintf("%g", a)
	case "BOOL":
		return v.value.(string)
	case "ARRAY":
		// interface{}[1,2,3,4]
		return v.value.(string)
	}
	fmt.Println("errortype", v)
	return "errortype"
}
func nestedPar(str string) []int {
	// (alo (blo (clo) gno) dno)
	depth := 0
	var arr []int
	for i := 0; i < len(str); i++ {
		if string(str[i]) == "(" {
			depth++
		} else if string(str[i]) == ")" {
			depth--
		}
		arr = append(arr, depth)
	}
	return arr
}
func getPart(arr []int, num int) (int, int) {
	// [1,1,1,2,2,2,6,6,6,3,3,3]
	start := -1
	end := 0
	for i := 0; i < len(arr); i++ {
		if start == -1 && arr[i] == num {
			start = i + 1
		}
		if start > 0 && arr[i+1] != num {
			end = i + 1
			break
		}
	}
	return start, end
}

func max(arr []int) int {
	if len(arr) <= 0 {
		return 0
	}
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

func replacePart[E any](arr []E, val E, start, end int) []E {
	// 1,2,3,4,5,6,7
	part1 := arr[:start]
	part2 := arr[end:]
	var newarr []E
	newarr = append(newarr, part1...)
	newarr = append(newarr, val)
	newarr = append(newarr, part2...)
	return newarr
}

func stringToFloat64(str string) float64 {
	r, err := strconv.ParseFloat(str, 64)
	if err != nil {
		panic(err)
	}
	return r
}

func determineType(varval string) string {
	if varval == "true" || varval == "false" {
		return "BOOL"
	} else if strings.Contains(varval, "[") {
		return "ARRAY"
	} else {
		return "NUMBER"
	}
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}
