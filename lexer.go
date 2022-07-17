package main

type tok struct {
	typ, value string
}

var signs = map[string]string{
	"\t": "TAB",
	" ":  "WHITESPACE",
	"\n": "NEWLINE",
	"+":  "PLUS",
	"-":  "MINUS",
	"*":  "STAR",
	"/":  "DASH",
	">":  "GREATER",
	"<":  "LESS",
	"=":  "EQUAL",
	"&":  "AND",
	"|":  "OR",
	"~":  "NOT",
	"{":  "LBRACKET",
	"}":  "RBRACKET",
	"(":  "LPAREN",
	")":  "RPAREN",
	"[":  "LSQUARE",
	"]":  "RSQUARE",
	";":  "SEMICOLON",
}

func Lex(str string) []tok {
	var token string = ""
	var tokens []tok
	str += "\n"
	for i := 0; i < len(str); i++ {
		s := string(str[i])
		if s == "#" {
			for string(str[i]) != "\n" {
				i++
			}
		} else if v, ok := signs[s]; ok {
			if token != "" {
				tokens = append(tokens, tok{"KEYWORD", token})
			}
			if v != "WHITESPACE" && v != "TAB" && v != "NEWLINE" {
				tokens = append(tokens, tok{v, s})
			}
			token = ""
		} else {
			token += s
		}
	}
	return tokens
}
