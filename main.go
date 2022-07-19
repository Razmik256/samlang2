package main

import (
	"fmt"
	"os"
)

func main() {
	filedata, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	fmt.Println(Interprate(Parse(Lex(string(filedata))), map[string]variable{}))
}
