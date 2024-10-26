package main

// import (
// 	"strings"
// )

// func codeGenerator(tokens []string, outputType string) string {
// 	points := map[string][]string{} // Maps variables names to point values
// 	tests := [][]string{}           // Maps test type to variable names

// 	// Parse tokens and stores points values and test calls
// 	for len(tokens) > 0 {
// 		// Stores points
// 		if tokens[0] == "ID" {
// 			varName := tokens[1]
// 			x := tokens[6]
// 			y := tokens[9]
// 			points[varName] = []string{x, y}
// 			tokens = tokens[11:]
// 		}
// 		// Stores test calls
// 		if tokens[0] == "TEST" {
// 			shape := tokens[2]
// 			testPoints := []string{shape}
// 			tokens = tokens[4:]
// 			// Gather all points in params
// 			for tokens[0] == "ID" {
// 				testPoints = append(testPoints, tokens[1])
// 				tokens = tokens[3:]
// 			}
// 			tests = append(tests, testPoints)
// 		}

// 		if tokens[0] == "PERIOD" {
// 			tokens = []string{}
// 		} else if tokens[0] == "SEMICOLON" {
// 			tokens = tokens[1:]
// 		}
// 	}
// 	// Generate code
// 	var code string

// 	// Output is dependent upon flag from outputType
// 	if outputType == "s" {
// 		// Iterate through tests
// 		for _, shape := range tests {
// 			code += "(process-" + strings.ToLower(shape[0])
// 			shape = shape[1:]
// 			// Iterate through params in test
// 			for _, param := range shape {
// 				code += " (make-point " + points[param][0] + " " + points[param][1] + ")"
// 			}
// 			code += ")\n"
// 		}
// 	} else {
// 		// Iterate through tests
// 		for _, shape := range tests {
// 			var code1, code2, testCode string
// 			shapeType := shape[0]
// 			if shapeType == "SQUARE" {
// 				testCode = "query(" + strings.ToLower(shape[0]) + "("
// 			} else {
// 				code1 = "query("
// 				code2 = "("
// 			}
// 			comment := "\n/* Processing test(" + strings.ToLower(shape[0])

// 			shape = shape[1:]
// 			// Iterate through params in test
// 			for _, param := range shape {
// 				comment += ", " + param
// 				if shapeType == "SQUARE" {
// 					testCode += "point2d(" + points[param][0] + ", " + points[param][1] + "), "
// 				} else {
// 					code2 += "point2d(" + points[param][0] + ", " + points[param][1] + "), "
// 				}

// 			}
// 			// Add comment to code output
// 			comment += ") */\n"
// 			code += comment

// 			// Create each query for when shape is a triangle
// 			if shapeType == "TRIANGLE" {
// 				code2 = code2[:len(code2)-2] + ")).\n"
// 				types := []string{
// 					"line",
// 					"triangle",
// 					"vertical",
// 					"horizontal",
// 					"equilateral",
// 					"isosceles",
// 					"right",
// 					"scalene",
// 					"acute",
// 					"obtuse",
// 				}
// 				for _, e := range types {
// 					code += code1 + e + code2
// 				}
// 			} else {
// 				testCode = testCode[:len(testCode)-2] + ")).\n"
// 				code += testCode
// 			}
// 		}
// 		// Query processing
// 		code += "\n/* Query Processing */\nwriteln(T) :- write(T), nl.\nmain:- forall(query(Q), Q-> (writeln(‘yes’)) ; (writeln(‘no’))),\nhalt."
// 	}

// 	return code

// }
