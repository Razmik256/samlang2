package main

import "fmt"

type coin struct {
	function, left, right string
}

func Parse(toks []tok) []coin {
	fmt.Println(toks, "\n===================")
	var coins []coin
	var ends []string
	for i := 0; i < len(toks); i++ {
		if toks[i].value == "}" {
			coins = append(coins, coin{ends[len(ends)-1], "", ""})
			ends = ends[:len(ends)-1]
		} else if toks[i].value == "=" && toks[i+1].typ != "=" {
			left := toks[i-1].value
			right := ""
			for toks[i+1].typ != "SEMICOLON" {
				i += 1
				right += toks[i].value
			}
			coins = append(coins, coin{"SET", left, right})
		} else if toks[i].value == "if" {
			val := ""
			for toks[i+1].typ != "LBRACKET" {
				i += 1
				val += toks[i].value
			}
			ends = append(ends, "IFEND")
			coins = append(coins, coin{"IF", val, ""})
		} else if toks[i].value == "else" {
			coins = append(coins, coin{"ELSE", "", ""})
			ends = append(ends, "ELSEEND")
		} else if toks[i].value == "func" {
			name, args := "", ""
			for toks[i+1].typ != "LPAREN" {
				i += 1
				name += toks[i].value
			}
			i += 1
			for toks[i+1].typ != "RPAREN" {
				i += 1
				args += toks[i].value
			}
			ends = append(ends, "FUNCEND")
			coins = append(coins, coin{"FUNC", name, args})
		}
	}
	return coins
}
