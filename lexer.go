package main

import (
	"fmt"
	"os"
	"regexp"
)

var lexemes = map[string]string{
	"point":    "POINT",
	";":        "SEMICOLON",
	",":        "COMMA",
	".":        "PERIOD",
	"(":        "LPAREN",
	")":        "RPAREN",
	"=":        "ASSIGN",
	"triangle": "TRIANGLE",
	"square":   "SQUARE",
	"test":     "TEST",
}

// Digit regexp
var numRegexp = regexp.MustCompile(`^[\d]+`)

// Identifier regexp
var idRegexp = regexp.MustCompile(`^[a-z]+`)

/*
*

	Lexical Analyzer
*/
func lexer(input string) []string {
	tokens := []string{}
	for len(input) > 0 {
		matched := false
		//Check for reserved words/lexemes
		for lexeme, token := range lexemes {
			if len(input) >= len(lexeme) && input[:len(lexeme)] == lexeme {
				tokens = append(tokens, token)
				input = input[len(lexeme):]
				matched = true
				break
			}
		}
		if matched {
			continue
		}
		//Checking digits, IDs and empty characters
		if lex := numRegexp.FindString(input); lex != "" { //Nums (0-9)
			tokens = append(tokens, "NUM", lex)
			input = input[len(lex):]
		} else if lex := idRegexp.FindString(input); lex != "" { //ID (a-z)
			tokens = append(tokens, "ID", lex)
			input = input[len(lex):]
		} else if input[0] == 32 { //Skip space character
			input = input[1:]
		} else if input[0] == 10 { //Skip new line character
			input = input[1:]
		} else { //Unkown Lexeme ERROR
			fmt.Println("Lexical error", input[:1], "not recognized")
			os.Exit(1)
		}

	}
	return tokens
}
