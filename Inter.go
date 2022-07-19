package main

import (
	"fmt"
	"regexp"
	"strconv"
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
	reg := regexp.MustCompile(`\w+\[\d[\,\d]*\]|\w+\[\d:\d\]`)
	for n, v := range vars {
		for strings.Contains(varval, n) && v.typ != "ARRAY" {
			varval = changePart(varval, n, v.value)
		}
	}
	if findstr := reg.FindString(varval); findstr != "" {
		parts := strings.Split(varval, "[")
		parts2 := strings.Split(parts[1], "]")
		parts[0] = parts[0][:len(parts[0])]
		parts[1] = parts2[0]
		// fmt.Println("p0", parts[0], "p1", parts[1])
		// y = x[1,2] -> x[1] ++ x[2];
		if arr, ok := vars[parts[0]]; ok {
			arr.value = arr.value[1 : len(arr.value)-1]
			var arrvars []float64
			for _, a := range strings.Split(arr.value, ",") {
				arrvars = append(arrvars, parseFloatNoErr(a))
			}
			if strings.Contains(parts[1], ",") {
				indexes := strings.Split(parts[1], ",")
				var values string
				values = "["
				for _, i := range indexes {
					ind, _ := strconv.Atoi(i)
					values += fmt.Sprintf("%g", arrvars[ind]) + ","
				}
				values = values[:len(values)-1]
				values += "]"
				varval = changePart(varval, findstr, values)
			} else if strings.Contains(parts[1], ":") {
				sides := strings.Split(parts[1], ":")
				side1, _ := strconv.Atoi(sides[0])
				side2, _ := strconv.Atoi(sides[1])
				var values string
				values = "["
				for i := side1; i < side2+1; i++ {
					values += fmt.Sprintf("%g", arrvars[i]) + ","
				}
				values = values[:len(values)-1]
				values += "]"
				varval = changePart(varval, findstr, values)
			} else {
				index, _ := strconv.Atoi(parts[1])
				varval = changePart(varval, findstr, fmt.Sprintf("%g", arrvars[index]))
				// fmt.Println("parts", parts[0], parts[1], "sides", side1, side2, "varval", varval, "findstr", findstr, "values", values)
			}
		} else {
			panic("No array with name " + parts[0])
		}
	}

	if strings.Contains(varval, "true") || strings.Contains(varval, "false") || strings.Contains(varval, ">") || strings.Contains(varval, "<") || strings.Contains(varval, "~") || strings.Contains(varval, "=") {
		vars[varname] = variable{"BOOL", BoolEval(varval, false)}
	} else if strings.Contains(varval, "[") || strings.Contains(varval, "++") {
		vars[varname] = variable{"ARRAY", ArrEval(varval)}
	} else {
		vars[varname] = variable{"NUMBER", fmt.Sprintf("%g", MathEval(varval, false))}
	}
}

func EvalComparison(str string, vars map[string]variable) bool {
	for n, v := range vars {
		for strings.Contains(str, n) {
			str = changePart(str, n, v.value)
		}
	}
	if BoolEval(str, false) == "true" {
		return true
	} else {
		return false
	}
}

func mergeVariables[vars map[string]variable](v1 vars, v2 vars) vars {
	for k, v := range v2 {
		v1[k] = v
	}
	return v1
}

func Interprate(cos []coin, protovars map[string]variable) map[string]variable {
	fmt.Println(cos, "\n==============")
	// whileends must be 1
	// time.Sleep(time.Second / 2)
	var vars = protovars

	for i := 0; i < len(cos); i++ {
		fmt.Println(cos[i])
		// contains " -> string (first, so it can't be "true" -> bool)
		// contains true/false -> bool
		// else -> number
		if cos[i].function == "SET" {
			EvalVar(cos[i].left, cos[i].right, vars)
		} else if cos[i].function == "IF" {
			if !EvalComparison(cos[i].left, vars) {
				for cos[i].function != "IFEND" {
					i++
				}
			}
		} else if cos[i].function == "WHILE" {
			if EvalComparison(cos[i].left, vars) {
				j := i
				whileends := 1
				for whileends != 0 {
					j++
					if cos[j].function == "WHILE" {
						whileends += 1
					} else if cos[j].function == "WHILEEND" {
						whileends -= 1
					}
				}
				for {
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
