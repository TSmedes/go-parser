package main

import (
	"strings"
)

func schemeGenerator(tokens []string) string {
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
	// Generate scheme code
	var schemeCode string
	// Iterate through tests
	for _, shape := range tests {
		schemeCode += "(process-" + strings.ToLower(shape[0])
		shape = shape[1:]
		// Iterate through params in test
		for _, param := range shape {
			schemeCode += " (make-point " + points[param][0] + " " + points[param][1] + ")"
		}
		schemeCode += ")\n"
	}

	return schemeCode

}
