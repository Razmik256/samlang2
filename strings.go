package main

import "fmt"

func stringIntoArr(str string) string {
	arr := "["
	str = str[1 : len(str)-1]
	for i := 0; i < len(str); i++ {
		arr += fmt.Sprintf("%d", str[i]) + ","
	}
	arr = arr[:len(arr)-1]
	arr += "]"
	return arr
}

func StringEval(str string) string {

	// "alo"

	return str
}
