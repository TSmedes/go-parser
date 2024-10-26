package main

import (
	"fmt"
	"os"
	"regexp"
)

/*
*

	Syntax Analyzer
*/
func syntax(tokens []string, outputType string) {

	tokens = stmtList(tokens)
	if len(tokens) > 0 {
		fmt.Println("Syntax error: Unknown symbols after program end .:", tokens[0])
		os.Exit(1)
		//ERROR
	}
}

func stmtList(tokens []string) []string {
	tokens = stmt(tokens)
	if tokens[0] == "PERIOD" {
		if len(tokens) > 1 {
			return tokens[1:]
		} else {
			return []string{}
		}
	}
	if tokens[0] == "SEMICOLON" {
		tokens = stmtList(tokens[1:])
	}
	return tokens
}

func stmt(tokens []string) []string {

	if tokens[0] == "ID" {
		tokens = pointDef(tokens)
	} else if tokens[0] == "TEST" {
		tokens = test(tokens)
	} else {
		fmt.Println("Syntax error: Invalid statement found:", tokens[0])
		os.Exit(1)
		//ERROR
	}
	return tokens
}

func pointDef(tokens []string) []string {
	// if tokens[0] != "ID" {
	// 	fmt.Println("Syntax error:", tokens[0], "found, ID expected")
	// 	os.Exit(1)
	// }
	// tokens = tokens[1:]

	tokens = id(tokens)

	if tokens[0] != "ASSIGN" {
		fmt.Println("Syntax error:", tokens[0], "found, = expected")
		os.Exit(1)
	}
	if tokens[1] != "POINT" {
		fmt.Println("Syntax error:", tokens[1], "found, point expected")
		os.Exit(1)
	}
	if tokens[2] != "LPAREN" {
		fmt.Println("Syntax error:", tokens[2], "found, ( expected")
		os.Exit(1)
	}
	tokens = tokens[3:]
	tokens = nums(tokens)

	if tokens[0] != "COMMA" {
		fmt.Println("Syntax error:", tokens[0], "found, , expected")
		os.Exit(1)
		//ERROR
	}

	tokens = tokens[1:]
	tokens = nums(tokens)

	if tokens[0] != "RPAREN" {
		fmt.Println("Syntax error:", tokens[0], "found, ) expected")
		os.Exit(1)
		//ERROR
	}
	return tokens[1:]
}

func test(tokens []string) []string {
	if tokens[0] != "TEST" {
		fmt.Println("Syntax error:", tokens[0], "found, test expected")
		os.Exit(1)
		//ERROR
	}
	if tokens[1] != "LPAREN" {
		fmt.Println("Syntax error:", tokens[1], "found, ( expected")
		os.Exit(1)
		//ERROR
	}

	tokens = tokens[2:]
	tokens = options(tokens)

	if tokens[0] != "COMMA" {
		fmt.Println("Syntax error:", tokens[0], "found, , expected")
		os.Exit(1)
		//ERROR
	}

	tokens = tokens[1:]
	tokens = pointList(tokens)
	if tokens[0] != "RPAREN" {
		fmt.Println("Syntax error:", tokens[0], "found, ) expected")
		os.Exit(1)
		//ERROR
	}

	return tokens[1:]
}

func id(tokens []string) []string {
	if tokens[0] != "ID" {
		fmt.Println("Syntax error: Invalid variable name found:", tokens[0])
	}
	alpha := regexp.MustCompile(`^[a-z]+$`).MatchString(tokens[1])
	if !alpha {
		fmt.Println("Syntax error: Invalid variable name found:", tokens[1])
		os.Exit(1)
		//ERROR
	}
	return tokens[2:]
}

func nums(tokens []string) []string {
	if tokens[0] != "NUM" {
		fmt.Println("Syntax error:", tokens[0], "found, integer expected")
		os.Exit(1)
	}
	numeric := regexp.MustCompile(`^[0-9]+$`).MatchString(tokens[1])
	if !numeric {
		fmt.Println("Syntax error:", tokens[1], "found, integer expected")
		os.Exit(1)
		//ERROR
	}
	return tokens[2:]
}

func options(tokens []string) []string {
	if tokens[0] != "TRIANGLE" && tokens[0] != "SQUARE" {
		fmt.Println("Syntax error:", tokens[0], "found, triangle or square expected")
		os.Exit(1)
		//ERROR
	}
	return tokens[1:]
}

func pointList(tokens []string) []string {
	tokens = id(tokens)
	if tokens[0] == "COMMA" {
		tokens = tokens[1:]
		tokens = pointList(tokens)
	}
	return tokens
}
