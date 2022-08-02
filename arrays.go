package main

// import (
// 	"fmt"
// 	"regexp"
// 	"strconv"
// 	"strings"
// )

// func ArrEval(varval string) string {
// 	// [1,2,3] ++ 1 -> [1,2,3,1]
// 	// 1 ++ 2 -> [1,2]
// 	// [1,2,3] ++ [1,2,3] -> [1,2,3,1,2,3]
// 	reg := regexp.MustCompile(`(\[\w+(\,\w+)*\]|\w+)\<\d+((\:\d+){0,1}|(\,\d+)*)\>`)

// 	for {
// 		// try to find array<id> parts
// 		findstr := reg.FindString(varval)
// 		if findstr == "" {
// 			break
// 		}
// 		parts := strings.Split(findstr, "<")
// 		parts2 := strings.Split(parts[1], ">")
// 		parts[0] = parts[0][1 : len(parts[0])-1] // [1,2,3]< -> 1,2,3
// 		parts[1] = parts2[0]                     // 1,2> -> 1,2
// 		var arrvals []string
// 		arrvals = append(arrvals, strings.Split(parts[0], ",")...)
// 		if strings.Contains(parts[1], ",") {
// 			// when [1,2,3,4]<1,3> -> 2, 4
// 			indexes := strings.Split(parts[1], ",")
// 			var values string
// 			values = "["
// 			for _, i := range indexes {
// 				ind, _ := strconv.Atoi(i)
// 				values += arrvals[ind] + ","
// 			}
// 			values = values[:len(values)-1]
// 			values += "]"
// 			varval = changePart(varval, findstr, values)
// 		} else if strings.Contains(parts[1], ":") {
// 			// [1,2,3,4,5]<1:3> -> 2,3,4
// 			sides := strings.Split(parts[1], ":")
// 			side1, _ := strconv.Atoi(sides[0])
// 			side2, _ := strconv.Atoi(sides[1])
// 			var values string
// 			values = "["
// 			for i := side1; i < side2+1; i++ {
// 				values += arrvals[i] + ","
// 			}
// 			values = values[:len(values)-1]
// 			values += "]"
// 			varval = changePart(varval, findstr, values)
// 		} else {
// 			// [1,2,3]<1> -> 2
// 			index, _ := strconv.Atoi(parts[1])
// 			varval = changePart(varval, findstr, arrvals[index])
// 		}
// 	}

// 	fmt.Println("arreval", varval)
// 	varval = strings.Replace(varval, "[]++", "", -1)
// 	varval = strings.Replace(varval, "++[]", "", -1)
// 	varval = strings.Replace(varval, "[", "", -1)
// 	varval = strings.Replace(varval, "]", "", -1)
// 	varval = strings.Replace(varval, "++", ",", -1)
// 	varval = "[" + varval + "]"
// 	return varval
// }
