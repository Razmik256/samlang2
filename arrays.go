package main

import (
	"fmt"
)

func ArrEval(val string) string {
	// [1,2,3] ++ 1 -> [1,2,3,1]
	// 1 ++ 2 -> [1,2]
	// [1,2,3] ++ [1,2,3]
	fmt.Println("arreval", val)
	// val = strings.Replace(val, "[", "", -1)
	// val = strings.Replace(val, "]", "", -1)
	// val = strings.Replace(val, ",", "++", -1)
	// parts := strings.Split(val, "++")
	// var arr string = "["
	// for _, v := range parts {
	// 	arr += v + ","
	// }
	// arr = arr[:len(arr)-1]
	// arr += "]"
	return val
}
