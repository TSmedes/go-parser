package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

/*
Main function
Purpose: this program accepts a file with 4Point grammar
and checks it lexically and syntactically. It then converts it
to either scheme or prolog code.
*/
func main() {
	// Check if the user has provided the correct number of arguments
	if len(os.Args) < 3 {
		fmt.Println("Missing parameter, usage:\ngo run . filename -flag\nflag can be p for prolog generation\nflag can be s for scheme generation")
		os.Exit(1)
	}
	if len(os.Args) > 3 {
		fmt.Println("Too many parameters, usage:\ngo run . filename -flag\nflag can be p for prolog generation\nflag can be s for scheme generation")
		os.Exit(1)
	}

	fileName := os.Args[1]
	outputType := os.Args[2]

	//Attempt opening file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("File not found")
		os.Exit(1)
	}
	//Check for correct flags
	if len(outputType) != 2 || outputType[0] != '-' || (outputType[1] != 'p' && outputType[1] != 's') {
		fmt.Println("Invalid flag, usage:\ngo run . filename -flag\nflag can be p for prolog generation\nflag can be s for scheme generation")
		os.Exit(1)
	}
	//Attempt reading file
	text, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file")
		os.Exit(1)
	}
	defer file.Close()

	//Lexical analysis
	tokens := lexer(string(text))

	//Syntactical Analysis
	syntax(tokens)

	//Adjust comments for scheme or prolog output
	if outputType[1] == 's' {
		fmt.Println("; processing input file", fileName)
		fmt.Println("; Lexical and Syntax analysis passed")
		fmt.Println("; Generating Scheme Code")
		fmt.Println(codeGenerator(tokens, "s"))
	} else {
		fmt.Println("/* processing input file", fileName)
		fmt.Println("   Lexical and Syntax analysis passed")
		fmt.Println("   Generating Prolog Code */")
		fmt.Println(codeGenerator(tokens, "p"))
	}

	os.Exit(0)
}

// Reserved words and characters and their respective token
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
	Purpose: The lexical analyzer parses teh input file and
	searches for valid characters and lexemes and generates a
	list of tokens
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

/*
*

	Syntax Analyzer
	purpose: the following functions
	are called recursively to analyze the program's
	syntax and ensure it follows the grammar rules
*/
func syntax(tokens []string) {

	tokens = stmtList(tokens)
	// Extra tokens after grammar is complete breaks the grammatical rules
	if len(tokens) > 0 {
		fmt.Println("Syntax error: Unknown symbols after program end .:", tokens[0])
		os.Exit(1)
	}
}

/*
Checks grammar for STMT_LIST
*/
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

/*
Checks grammar for STMT
*/
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

/*
Checks grammar for POINT_DEF
*/
func pointDef(tokens []string) []string {
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

/*
Checks grammar for TEST
*/
func test(tokens []string) []string {
	if tokens[0] != "TEST" {
		fmt.Println("Syntax error:", tokens[0], "found, test expected")
		os.Exit(1)
	}
	if tokens[1] != "LPAREN" {
		fmt.Println("Syntax error:", tokens[1], "found, ( expected")
		os.Exit(1)
	}

	tokens = tokens[2:]
	tokens = options(tokens)

	if tokens[0] != "COMMA" {
		fmt.Println("Syntax error:", tokens[0], "found, , expected")
		os.Exit(1)
	}

	tokens = tokens[1:]
	tokens = pointList(tokens)
	if tokens[0] != "RPAREN" {
		fmt.Println("Syntax error:", tokens[0], "found, ) expected")
		os.Exit(1)
	}

	return tokens[1:]
}

/*
Checks grammar for ID
*/
func id(tokens []string) []string {
	if tokens[0] != "ID" {
		fmt.Println("Syntax error: Invalid variable name found:", tokens[0])
	}
	alpha := regexp.MustCompile(`^[a-z]+$`).MatchString(tokens[1])
	if !alpha {
		fmt.Println("Syntax error: Invalid variable name found:", tokens[1])
		os.Exit(1)
	}
	return tokens[2:]
}

/*
Checks grammar for NUMS
*/
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

/*
Checks grammar for OPTION
*/
func options(tokens []string) []string {
	if tokens[0] != "TRIANGLE" && tokens[0] != "SQUARE" {
		fmt.Println("Syntax error:", tokens[0], "found, triangle or square expected")
		os.Exit(1)
	}
	return tokens[1:]
}

/*
Checks grammar for POINT_LIST
*/
func pointList(tokens []string) []string {
	tokens = id(tokens)
	if tokens[0] == "COMMA" {
		tokens = tokens[1:]
		tokens = pointList(tokens)
	}
	return tokens
}

/*
*

	codeGenerator
	purpose: Generates the output code in the desired language
	from the input flag (scheme or prolog)
*/
func codeGenerator(tokens []string, outputType string) string {
	points := map[string][]string{} // Maps variables names to point values
	tests := [][]string{}           // Maps test type to variable names

	// Parse tokens and stores points values and test calls
	for len(tokens) > 0 {
		// Stores points
		if tokens[0] == "ID" {
			varName := tokens[1]
			x := tokens[6]
			y := tokens[9]
			points[varName] = []string{x, y}
			tokens = tokens[11:]
		}
		// Stores test calls
		if tokens[0] == "TEST" {
			shape := tokens[2]
			testPoints := []string{shape}
			tokens = tokens[4:]
			// Gather all points in params
			for tokens[0] == "ID" {
				testPoints = append(testPoints, tokens[1])
				tokens = tokens[3:]
			}
			tests = append(tests, testPoints)
		}

		if tokens[0] == "PERIOD" {
			tokens = []string{}
		} else if tokens[0] == "SEMICOLON" {
			tokens = tokens[1:]
		}
	}
	// Generate code
	var code string

	// Output is dependent upon flag from outputType
	if outputType == "s" {
		// Iterate through tests
		for _, shape := range tests {
			code += "(process-" + strings.ToLower(shape[0])
			shape = shape[1:]
			// Iterate through params in test
			for _, param := range shape {
				code += " (make-point " + points[param][0] + " " + points[param][1] + ")"
			}
			code += ")\n"
		}
	} else {
		// Iterate through tests
		for _, shape := range tests {
			var code1, code2, testCode string
			shapeType := shape[0]
			if shapeType == "SQUARE" {
				testCode = "query(" + strings.ToLower(shape[0]) + "("
			} else {
				code1 = "query("
				code2 = "("
			}
			comment := "\n/* Processing test(" + strings.ToLower(shape[0])

			shape = shape[1:]
			// Iterate through params in test
			for _, param := range shape {
				comment += ", " + param
				if shapeType == "SQUARE" {
					testCode += "point2d(" + points[param][0] + ", " + points[param][1] + "), "
				} else {
					code2 += "point2d(" + points[param][0] + ", " + points[param][1] + "), "
				}

			}
			// Add comment to code output
			comment += ") */\n"
			code += comment

			// Create each query for when shape is a triangle
			if shapeType == "TRIANGLE" {
				code2 = code2[:len(code2)-2] + ")).\n"
				types := []string{
					"line",
					"triangle",
					"vertical",
					"horizontal",
					"equilateral",
					"isosceles",
					"right",
					"scalene",
					"acute",
					"obtuse",
				}
				for _, e := range types {
					code += code1 + e + code2
				}
			} else {
				testCode = testCode[:len(testCode)-2] + ")).\n"
				code += testCode
			}
		}
		// Query processing
		code += "\n/* Query Processing */\nwriteln(T) :- write(T), nl.\nmain:- forall(query(Q), Q-> (writeln(‘yes’)) ; (writeln(‘no’))),\nhalt."
	}

	return code

}
