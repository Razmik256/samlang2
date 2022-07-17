package main

import (
	"fmt"
	"strings"
)

type variable struct {
	typ, value string
}

func changePart(str, change, value string) string {
	start := strings.Index(str, change)
	end := start + len(change)
	return replacePart(str, value, start, end)
}

func EvalVar(varname, varval string, vars map[string]variable) {
	for n, v := range vars {
		if strings.Contains(varval, n) {
			varval = changePart(varval, n, v.value)
		}
	}
	vars[varname] = variable{"NUMBER", fmt.Sprintf("%f", MathEval(varval, false))}
}

func Interprate(cos []coin) map[string]variable {
	fmt.Println(cos, "\n==============")

	var vars = map[string]variable{}

	for i := 0; i < len(cos); i++ {
		// contains " -> string (first, so it can't be "true" -> bool)
		// contains true/false -> bool
		// else -> number
		if cos[i].function == "SET" {
			if strings.Contains(cos[i].right, "true") || strings.Contains(cos[i].right, "false") {
				vars[cos[i].left] = variable{"BOOL", fmt.Sprintf("%t", BoolEval(cos[i].right, false))}
			} else {
				EvalVar(cos[i].left, cos[i].right, vars)
			}
		}
	}

	return vars
}
