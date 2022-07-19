package main

import (
	"fmt"
	"strings"
)

type variable struct {
	typ, value string
}

func changePart(str, change, value string) string {
	// changes string by part
	start := strings.Index(str, change)
	end := start + len(change)
	return replacePart(str, value, start, end)
}

func EvalVar(varname, varval string, vars map[string]variable) {
	// evalueates any given value (numbers, booleans, strings, arrays) in string form
	// aply variables into string
	for n, v := range vars {
		for strings.Contains(varval, n) {
			varval = changePart(varval, n, v.value)
		}
	}

	if strings.Contains(varval, "[") || strings.Contains(varval, "++") {
		vars[varname] = variable{"ARRAY", ArrEval(varval)}
	} else if strings.Contains(varval, "true") || strings.Contains(varval, "false") || strings.Contains(varval, ">") || strings.Contains(varval, "<") || strings.Contains(varval, "~") || strings.Contains(varval, "=") {
		vars[varname] = variable{"BOOL", BoolEval(varval, false)}
	} else {
		vars[varname] = variable{"NUMBER", fmt.Sprintf("%g", MathEval(varval, false))}
	}
}

func EvalComparison(str string, vars map[string]variable) bool {
	// aply variables
	for n, v := range vars {
		for strings.Contains(str, n) {
			str = changePart(str, n, v.value)
		}
	}
	// make comparison
	if BoolEval(str, false) == "true" {
		return true
	} else {
		return false
	}
}

// merges new (local) variables with old ones
func mergeVariables[vars map[string]variable](v1 vars, v2 vars) vars {
	for k, v := range v2 {
		v1[k] = v
	}
	return v1
}

func Interprate(cos []coin, protovars map[string]variable) map[string]variable {
	fmt.Println(cos, "\n==============")
	// time.Sleep(time.Second / 2)
	var vars = protovars

	for i := 0; i < len(cos); i++ {
		fmt.Println(cos[i])
		// contains " -> string (first, so it can't be "true" -> bool)
		// contains [ -> array
		// contains true/false -> bool
		// else -> number
		if cos[i].function == "SET" {
			EvalVar(cos[i].left, cos[i].right, vars)
		} else if cos[i].function == "IF" {
			end := i
			beforeelse := i
			ifends := 1
			for ifends > 0 {
				// find last ifend
				end++
				if cos[end].function == "IFEND" {
					ifends--
				} else if cos[end].function == "IF" {
					ifends++
				}
			}
			// find last else
			beforeelse = end
			if cos[end+1].function == "ELSE" {
				for cos[end].function != "ELSEEND" {
					end++
				}
			}
			// if true, interprate it and go to end
			if EvalComparison(cos[i].left, vars) {
				vars = mergeVariables(vars, Interprate(cos[i+1:beforeelse], vars))
				i = end
			} else {
				// if not, go to else
				i = beforeelse + 1
			}
		} else if cos[i].function == "WHILE" {
			if EvalComparison(cos[i].left, vars) {
				j := i
				whileends := 1
				for whileends != 0 {
					// find last while
					j++
					if cos[j].function == "WHILE" {
						whileends += 1
					} else if cos[j].function == "WHILEEND" {
						whileends -= 1
					}
				}
				for {
					// while it's true, interprate the part inside while
					if EvalComparison(cos[i].left, vars) {
						vars = mergeVariables(vars, Interprate(cos[i+1:j], vars))
					} else {
						break
					}
				}
				i = j
			} else {
				for cos[i].function != "WHILEEND" {
					i++
				}
			}
		}
	}

	return vars
}
