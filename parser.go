package main

import (
	"fmt"
	"io"
	"os"
)

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
