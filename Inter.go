package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

type variable struct {
	typ   string
	value interface{}
}
type function struct {
	args       []string
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

func MurderFunction(f string) (name string, args []string) {
	parts := strings.Split(f, "(")
	name = parts[0]
	token := ""
	for i := 0; i < len(parts[1]); i++ {
		a := parts[1]
		if string(a[i]) == "[" {
			j := i
			for string(a[i]) != "]" {
				i++
			}
			args = append(args, a[j:i+1])
			token = ""
		} else if string(a[i]) == "," {
			if token != "" {
				args = append(args, token)
				token = ""
			}
		} else if string(a[i]) == ")" {
			args = append(args, token)
			token = ""
		} else {
			token += string(a[i])
		}
	}
	return name, args
}

func funcEval(f function, args []string, cos []coin, vars map[string]variable, funcs map[string]function) variable {
	funcvars := make(map[string]variable)
	for k, v := range vars {
		funcvars[k] = v
	}
	for _, v := range args {
		fmt.Println("whereisargs", v)
	}
	// problem, where is second variable
	for i, v := range f.args {
		fmt.Println("arg", i, "v", v, "args[i]", args)
		funcvars[v] = EvalValue(setVariablesInVal(args[i], cos, vars, funcs))
	}
	fmt.Println("funcEval", f, "funcvars", funcvars, "args", args, "vars", vars)
	ret := Interprate(cos[f.start:f.end], funcvars, funcs)["return"]
	return ret
}

func setVariablesInVal(value string, cos []coin, vars map[string]variable, funcs map[string]function) string {

	// for n, v := range vars {
	// 	// ALO
	// 	// I HAVE PROBLEM WITH VARIABLE NAMES, arr = 0; newarr -> new0

	// 	// x = alo;
	// 	// x = alo();
	// 	if strings.Contains(value, n) {

	// 		// i should check it's not a function
	// 		value = strings.ReplaceAll(value, n, interfaceIntoString(v))
	// 	}
	// }
	l := Lex(value)
	for i, v := range l {
		if val, ok := vars[v.value]; ok {
			if i+1 < len(l) {
				if l[i+1].value != "(" {
					l[i] = tok{"KEYWORD", val.value.(string)}
				}
			} else {
				l[i] = tok{"KEYWORD", val.value.(string)}
			}
		}
	}
	value = unLex(l)

	funcreg := regexp.MustCompile(`\w+\(.*\)`)
	for _, f := range funcreg.FindAllString(value, -1) {
		name, args := MurderFunction(f)
		if thefunc, ok := funcs[name]; ok {
			value = funcreg.ReplaceAllString(value, interfaceIntoString(funcEval(thefunc, args, cos, vars, funcs)))
		} else {
			panic("no such function : " + f)
		}
	}

	return value
}

func EvalValue(varval string) variable {
	// all variables must be already settled in varval
	// this will eval only varvals like 2 + 2, true & false, [1,2,3] ++ 1
	// assume we have something like ~(2 + 3/[1,2,3]<2> < 6*[7,5]<1:1>) & true
	// can I run this? lets see. P.S. apparently I can :D
	for {
		depth := nestedPar(varval)
		m := max(depth)
		if m == 0 {
			break
		}
		s1, s2 := getPart(depth, m)
		fmt.Println(varval)
		varval = strings.Replace(varval, varval[s1-1:s2+1], EvalValueHelper(varval[s1:s2]), -1)
	}

	val := EvalValueHelper(varval)
	fmt.Println("varval", varval, "val", val)
	return variable{determineType(val), val}
}

type mathLex struct {
	strval  string
	realval float64
	issign  bool
	typ     string
}

func getArrayValues(truLex []mathLex, end int) (int, []mathLex) {
	// [ 1 , 2 , 3 , 4 , [ 1 , 2 , 3 ] ]
	// loop goes only to left
	if truLex[end].typ == "SIGN" {
		squares := 1
		j := end
		for squares != 0 {
			j--
			if truLex[j].strval == "]" {
				squares++
			} else if truLex[j].strval == "[" {
				squares--

			}
		}
		start := j
		var inArrVals []mathLex
		if truLex[j+1].issign == false || truLex[j+1].typ == "BOOL" {
			inArrVals = append(inArrVals, truLex[j+1])
			// add the first value if it's not [
		}
		for j := start + 1; j < end; j++ {
			if truLex[j].strval == "]" {
				break
			} else if truLex[j].strval == "[" {
				var val []mathLex
				val = append(val, truLex[j])
				for truLex[j].strval != "]" {
					j++
					val = append(val, truLex[j])
				}
				var strval string
				for _, v := range val {
					strval += v.strval
				}
				inArrVals = append(inArrVals, mathLex{strval, 0, false, "ARRAY"})
			} else if truLex[j].strval == "," && truLex[j+1].strval != "[" {
				inArrVals = append(inArrVals, truLex[j+1])
			}

		}
		return start, inArrVals

	} else {
		// string [1,2,3]
		strarr := strings.Split(truLex[end].strval, ",")
		strarr[0] = strarr[0][1:]
		l := len(strarr) - 1
		strarr[l] = strarr[l][:len(strarr[l])-1] // wtf (i delete the last rune of the last value)
		var numbers []mathLex
		for _, v := range strarr {
			numbers = append(numbers, mathLex{v, stringToFloat64(v), false, determineType(v)})
		}
		return end, numbers
	}
}

func lexToTruLex(varval string) []mathLex {
	var truLex []mathLex
	for _, v := range Lex(varval) {
		a, err := strconv.ParseFloat(v.value, 64)
		if err == nil {
			truLex = append(truLex, mathLex{v.value, a, false, "NUMBER"})
		} else if v.value == "true" || v.value == "false" {
			truLex = append(truLex, mathLex{v.value, 0, true, "BOOL"})
		} else {
			truLex = append(truLex, mathLex{v.value, 0, true, "SIGN"})
		}
	}
	return truLex
}

func EvalValueHelper(varval string) string {
	// no we have to walk trought operators (math + - / *, boolean ~ & | < > ==, arrays ^ !!x)
	// + MATH
	// + BOOL
	// - ARRAY
	// - STRINGS

	// because I do this in one single function, maybe static typing can improve performance, but fuck that I'm too lazy

	// MATH LETS GO DUDE
	var truLex []mathLex

	truLex = lexToTruLex(varval)

	// [1,2,3]^[1,2,3];
	for {
		if i := slices.Index(truLex, mathLex{"^", 0, true, "SIGN"}); i != -1 {
			// for first I have to find the end of the second array
			// then get the two array values by function
			// and by iteration append them to each other
			if truLex[i+1].issign == false || truLex[i+1].typ == "BOOL" {
				start, arr1 := getArrayValues(truLex, i-1)
				fmt.Println("arr1", truLex[i-1])
				var newarr []mathLex
				newarr = append(newarr, arr1...)
				newarr = append(newarr, truLex[i+1])
				var stringvar string
				stringvar += "["
				for _, v := range newarr {
					stringvar += v.strval + ","
				}
				stringvar = stringvar[:len(stringvar)-1]
				stringvar += "]"
				truLex = replacePart(truLex, mathLex{stringvar, 0, false, "ARRAY"}, start, i+2)
			} else {
				squares := 1
				j := i + 1
				for squares != 0 {
					j++
					if truLex[j].strval == "[" {
						squares++
					} else if truLex[j].strval == "]" {
						squares--
					}
				}
				start, arr1 := getArrayValues(truLex, i-1)
				_, arr2 := getArrayValues(truLex, j)
				var newarr []mathLex
				newarr = append(newarr, arr1...)
				newarr = append(newarr, arr2...)
				var stringvar string
				stringvar += "["
				for _, v := range newarr {
					stringvar += v.strval + ","
				}
				stringvar = stringvar[:len(stringvar)-1]
				stringvar += "]"
				truLex = replacePart(truLex, mathLex{stringvar, 0, false, "ARRAY"}, start, j+1)
			}

		} else {
			break
		}
	}

	// array indexing
	for {
		if i := slices.Index(truLex, mathLex{"!!", 0, true, "SIGN"}); i != -1 {
			fmt.Println(truLex[i-1])
			if truLex[i-1].issign == true || truLex[i-1].typ == "ARRAY" {

				index := int(truLex[i+1].realval)
				bangval := mathLex{"error", 0, false, "error"}
				start, arr := getArrayValues(truLex, i-1)
				bangval = arr[index]
				if bangval.typ == "NUMBER" {
					truLex = replacePart(truLex, mathLex{bangval.strval, bangval.realval, false, "NUMBER"}, start, i+2)
				} else if bangval.typ == "BOOL" {
					truLex = replacePart(truLex, mathLex{bangval.strval, bangval.realval, false, "BOOL"}, start, i+2)
				} else if bangval.typ == "ARRAY" {
					truLex = replacePart(truLex, mathLex{bangval.strval, bangval.realval, false, "ARRAY"}, start, i+2)
				}
			} else {
				panic("wha iz dis? " + truLex[i-1].strval)
			}
		} else {
			break
		}
	}

	for {
		if i := slices.Index(truLex, mathLex{"-", 0, true, "SIGN"}); i != -1 {
			if i-1 >= 0 {
				if truLex[i-1].issign == true {
					truLex = replacePart(truLex, mathLex{truLex[i+1].strval, -truLex[i+1].realval, false, "NUMBER"}, i, i+2)
				} else {
					// I dunno why does this work but it works :P
					break
				}
			} else {
				// esincha yanm chjoga
				truLex = replacePart(truLex, mathLex{truLex[i+1].strval, -truLex[i+1].realval, false, "NUMBER"}, i, i+2)
			}
		} else {
			break
		}
	}
	for {
		if i := slices.Index(truLex, mathLex{"*", 0, true, "SIGN"}); i != -1 {
			a := truLex[i-1].realval * truLex[i+1].realval
			truLex = replacePart(truLex, mathLex{fmt.Sprintf("%g", a), a, false, "NUMBER"}, i-1, i+2)
		} else {
			break
		}
	}
	for {
		if i := slices.Index(truLex, mathLex{"/", 0, true, "SIGN"}); i != -1 {
			a := truLex[i-1].realval / truLex[i+1].realval
			truLex = replacePart(truLex, mathLex{fmt.Sprintf("%g", a), a, false, "NUMBER"}, i-1, i+2)
		} else {
			break
		}
	}
	for {
		if i := slices.Index(truLex, mathLex{"-", 0, true, "SIGN"}); i != -1 {
			a := truLex[i-1].realval - truLex[i+1].realval
			truLex = replacePart(truLex, mathLex{fmt.Sprintf("%g", a), a, false, "NUMBER"}, i-1, i+2)
		} else {
			break
		}
	}
	for {
		if i := slices.Index(truLex, mathLex{"+", 0, true, "SIGN"}); i != -1 {
			a := truLex[i-1].realval + truLex[i+1].realval
			truLex = replacePart(truLex, mathLex{fmt.Sprintf("%g", a), a, false, "NUMBER"}, i-1, i+2)
		} else {
			break
		}
	}

	// COMPARISON
	for i := 0; i < len(truLex); i++ {
		if truLex[i].strval == "==" {
			truLex = replacePart(truLex, mathLex{fmt.Sprintf("%t", truLex[i-1].realval == truLex[i+1].realval), 0, true, "BOOL"}, i-1, i+2)
		} else if truLex[i].strval == "<" {
			truLex = replacePart(truLex, mathLex{fmt.Sprintf("%t", truLex[i-1].realval < truLex[i+1].realval), 0, true, "BOOL"}, i-1, i+2)
		} else if truLex[i].strval == ">" {
			truLex = replacePart(truLex, mathLex{fmt.Sprintf("%t", truLex[i-1].realval > truLex[i+1].realval), 0, true, "BOOL"}, i-1, i+2)
		} else if truLex[i].strval == "~=" {
			truLex = replacePart(truLex, mathLex{fmt.Sprintf("%t", truLex[i-1].realval != truLex[i+1].realval), 0, true, "BOOL"}, i-1, i+2)
		}
	}
	// BOOLS
	for i := 0; i < len(truLex); i++ {
		if truLex[i].strval == "&" {
			v1, err := strconv.ParseBool(truLex[i-1].strval)
			Check(err)
			v2, err := strconv.ParseBool(truLex[i+1].strval)
			Check(err)
			truLex = replacePart(truLex, mathLex{fmt.Sprintf("%t", v1 && v2), 0, true, "BOOL"}, i-1, i+2)
		} else if truLex[i].strval == "|" {
			v1, err := strconv.ParseBool(truLex[i-1].strval)
			Check(err)
			v2, err := strconv.ParseBool(truLex[i+1].strval)
			Check(err)
			truLex = replacePart(truLex, mathLex{fmt.Sprintf("%t", v1 || v2), 0, true, "BOOL"}, i-1, i+2)
		} else if truLex[i].strval == "~" {
			val, err := strconv.ParseBool(truLex[i+1].strval)
			Check(err)
			truLex = replacePart(truLex, mathLex{fmt.Sprintf("%t", !val), 0, true, "BOOL"}, i, i+2)
		}
	}

	newvarval := ""
	for _, v := range truLex {
		newvarval += v.strval
	}
	return newvarval
}

// should i send the functions too?...
func Interprate(cos []coin, vars map[string]variable, funcs map[string]function) map[string]variable {
	fmt.Println(cos, "\n===================")
	for i := 0; i < len(cos); i++ {
		if cos[i].function == "SET" {
			vars[cos[i].left] = EvalValue(setVariablesInVal(cos[i].right, cos, vars, funcs))
		} else if cos[i].function == "IF" {
			e := EvalValue(setVariablesInVal(cos[i].left, cos, vars, funcs)).value
			if e != "true" && e != "false" {
				panic("wrong if statement")
			}
			if e == "true" {
				j := i
				for cos[i].function != "IFEND" {
					i++
				}
				a := Interprate(cos[j+1:i], vars, funcs)
				if len(a) > 0 {
					mergeVariables(vars, a)
				}
				if cos[i+1].function == "ELSE" {
					for cos[i].function != "ELSEEND" {
						i++
					}
				}
			} else {
				for cos[i].function != "IFEND" {
					i++
				}
				j := i
				for cos[i].function != "ELSEEND" {
					i++
				}
				a := Interprate(cos[j+1:i], vars, funcs)
				if len(a) > 0 {
					mergeVariables(vars, a)
				}
			}
		} else if cos[i].function == "WHILE" {
			j := i
			for cos[i].function != "WHILEEND" {
				i++
			}
			for EvalValue(setVariablesInVal(cos[j].left, cos, vars, funcs)).value == "true" {
				a := Interprate(cos[j+1:i], vars, funcs)
				if len(a) > 0 {
					mergeVariables(vars, a)
				}
			}
		} else if cos[i].function == "FUNC" {
			fmt.Println(cos[i])
			args := strings.Split(cos[i].right, ",")
			j := i
			for cos[i].function != "FUNCEND" {
				i++
			}
			funcs[cos[j].left] = function{args, j + 1, i}
		} else if cos[i].function == "RETURN" {
			return map[string]variable{"return": EvalValue(setVariablesInVal(cos[i].left, cos, vars, funcs))}
		}
	}
	// } else if cos[i].function == "IF" {
	// 	v := EvalValue(setVariablesInVal(cos[i].right, vars, funcs))
	// 	if v.typ == "BOOL" {
	// 		if v.value == "true" {

	// 		} else {

	// 		}
	// 	} else {
	// 		panic("wrong comparison " + cos[i].left)
	// 	}
	// }
	return vars
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

// func Interprate(cos []coin, protovars map[string]variable, protofuncs map[string]function) map[string]variable {
// 	fmt.Println(cos, "\n==============")
// 	// time.Sleep(time.Second / 2)
// 	var vars = protovars
// 	var funcs = protofuncs

// 	for i := 0; i < len(cos); i++ {
// 		fmt.Println(cos[i])
// 		if cos[i].function == "SET" {
// 			EvalVar(cos[i].left, cos[i].right, vars, funcs)
// 			vars[cos[i].left] = EvalValue(setVariables(cos[i].right, vars, funcs))
// 		} else if cos[i].function == "IF" {
// 			end := i
// 			beforeelse := i
// 			ifends := 1
// 			for ifends > 0 {
// 				// find last ifend
// 				end++
// 				if cos[end].function == "IFEND" {
// 					ifends--
// 				} else if cos[end].function == "IF" {
// 					ifends++
// 				}
// 			}
// 			// find last else
// 			beforeelse = end
// 			if cos[end+1].function == "ELSE" {
// 				for cos[end].function != "ELSEEND" {
// 					end++
// 				}
// 			}
// 			// if true, interprate it and go to end
// 			if EvalComparison(cos[i].left, vars) {
// 				mergeVariables(vars, Interprate(cos[i+1:beforeelse], vars, funcs))
// 				i = end
// 			} else {
// 				// if not, go to else
// 				i = beforeelse + 1
// 			}
// 		} else if cos[i].function == "WHILE" {
// 			if EvalComparison(cos[i].left, vars) {
// 				j := i
// 				whileends := 1
// 				for whileends != 0 {
// 					// find last while
// 					j++
// 					if cos[j].function == "WHILE" {
// 						whileends += 1
// 					} else if cos[j].function == "WHILEEND" {
// 						whileends -= 1
// 					}
// 				}
// 				for {
// 					// while it's true, interprate the part inside while
// 					if EvalComparison(cos[i].left, vars) {
// 						mergeVariables(vars, Interprate(cos[i+1:j], vars, funcs))
// 					} else {
// 						break
// 					}
// 				}
// 				i = j
// 			} else {
// 				for cos[i].function != "WHILEEND" {
// 					i++
// 				}
// 			}
// 		} else if cos[i].function == "FUNC" {
// 			start := i
// 			for cos[i].function != "FUNCEND" {
// 				i++
// 			}
// 			funcs[cos[start].left] = function{cos[start].right, start, i}
// 		} else if cos[i].function == "CALL" {
// 			if f, ok := funcs[cos[i].left]; ok {
// 				var localvars = map[string]variable{}
// 				funcargs := strings.Split(f.args, ",")
// 				for k, arg := range strings.Split(cos[i].right, ",") {
// 					localvars[funcargs[k]] = returnTypedVariable(arg)
// 				}
// 				mergeVariables(localvars, vars)
// 				Interprate(cos[f.start+1:f.end], localvars, funcs)
// 			}
// 		}
// 	}

// 	fmt.Println("funcs", funcs)
// 	return vars
// }
