package main

import (
	"fmt"
	"regexp"
	"strings"
)

type variable struct {
	typ   string
	value interface{}
}
type function struct {
	args       string
	start, end int
}

// func changePart(str, change, value string) string {
// 	// changes string by part
// 	start := strings.Index(str, change)
// 	end := start + len(change)
// 	return replacePart(str, value, start, end)
// }

// func returnTypedVariable(varval string) variable {
// 	if strings.Contains(varval, "\"") {
// 		return variable{"STRING", stringIntoArr(varval)}
// 	} else if strings.Contains(varval, "[") || strings.Contains(varval, "++") {
// 		return variable{"ARRAY", ArrEval(varval)}
// 	} else if strings.Contains(varval, "true") || strings.Contains(varval, "false") || strings.Contains(varval, ">") || strings.Contains(varval, "<") || strings.Contains(varval, "~") || strings.Contains(varval, "=") {
// 		return variable{"BOOL", BoolEval(varval, false)}
// 	} else {
// 		return variable{"NUMBER", fmt.Sprintf("%g", MathEval(varval, false))}
// 	}
// }

// func EvalVar(varname, varval string, vars map[string]variable, funcs map[string]function) {
// 	// evalueates any given value (numbers, booleans, strings, arrays) in string form
// 	// aply variables into string
// 	for n, v := range vars {
// 		for strings.Contains(varval, n) {
// 			varval = changePart(varval, n, v.value)
// 		}
// 	}

// 	// contains " -> string (first, so it can't be "true" -> bool)
// 	// contains [ -> array
// 	// contains true/false -> bool
// 	// else -> number
// 	funcreg := regexp.MustCompile(`\w+\(\w+(\,\w+)*\)`)
// 	if funcreg.MatchString(varval) {
// 		funcCoin := Parse(Lex("func " + varval))[0]
// 		if f, ok := funcs[funcCoin.left]; ok {
// 			vars[varname] = funcEval(f, strings.Split(funcCoin.right, ","))
// 		} else {
// 			panic("no such function " + funcCoin.left)
// 		}
// 	} else {
// 		vars[varname] = returnTypedVariable(varval)
// 	}
// }

func interfaceIntoString(v variable) string {
	switch v.typ {
	case "number":
		return fmt.Sprintf("%g", v.value)
	case "bool":
		return fmt.Sprintf("%t", v.value)
	case "array":
		// interface{}[1,2,3,4]
		return "not implemented, i dunno how does it look"
	}
	return "error type"
}

func setVariables(value string, vars map[string]variable, funcs map[string]function) string {
	for n, v := range vars {
		if strings.Contains(value, n) {
			strings.ReplaceAll(value, n, interfaceIntoString(v))
		}
	}
	funcreg := regexp.MustCompile(`\w+\(\w+(\,\w+)*\)`)
	for _, f := range funcreg.FindAllString(value, -1) {
		name, args := MurderFunction(f)
		if thefunc, ok := funcs[f]; ok {
			funcreg.ReplaceAllString(value, interfaceIntoString(funcEval(thefunc, args)))
		} else {
			panic("no such function : " + f)
		}
	}
	return value
}

// func EvalComparison(str string, vars map[string]variable) bool {
// 	// aply variables
// 	for n, v := range vars {
// 		for strings.Contains(str, n) {
// 			str = changePart(str, n, v.value)
// 		}
// 	}
// 	// make comparison
// 	if BoolEval(str, false) == "true" {
// 		return true
// 	} else {
// 		return false
// 	}
// }

// merges new (local) variables with old ones
func mergeVariables[vars map[string]variable](v1 vars, v2 vars) {
	for k, v := range v2 {
		v1[k] = v
	}
}

func Interprate(cos []coin, protovars map[string]variable, protofuncs map[string]function) map[string]variable {
	fmt.Println(cos, "\n==============")
	// time.Sleep(time.Second / 2)
	var vars = protovars
	var funcs = protofuncs

	for i := 0; i < len(cos); i++ {
		fmt.Println(cos[i])
		if cos[i].function == "SET" {
			EvalVar(cos[i].left, cos[i].right, vars, funcs)
			vars[cos[i].left] = EvalValue(setVariables(cos[i].right, vars, funcs))
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
				mergeVariables(vars, Interprate(cos[i+1:beforeelse], vars, funcs))
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
						mergeVariables(vars, Interprate(cos[i+1:j], vars, funcs))
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
		} else if cos[i].function == "FUNC" {
			start := i
			for cos[i].function != "FUNCEND" {
				i++
			}
			funcs[cos[start].left] = function{cos[start].right, start, i}
		} else if cos[i].function == "CALL" {
			if f, ok := funcs[cos[i].left]; ok {
				var localvars = map[string]variable{}
				funcargs := strings.Split(f.args, ",")
				for k, arg := range strings.Split(cos[i].right, ",") {
					localvars[funcargs[k]] = returnTypedVariable(arg)
				}
				mergeVariables(localvars, vars)
				Interprate(cos[f.start+1:f.end], localvars, funcs)
			}
		}
	}

	fmt.Println("funcs", funcs)
	return vars
}
