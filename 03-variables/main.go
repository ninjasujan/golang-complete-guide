package main

import (
	"fmt"
	"strconv"
)

// main demonstrates various string and numeric conversions in Go.
// It covers the following topics:
//   - Parsing a string to an integer using strconv.Atoi, including error handling for invalid input.
//   - Printing the parsed integer and its type.
//   - Parsing a string to a boolean using strconv.ParseBool, with error handling for invalid boolean strings.
//   - Converting an integer to its string representation using strconv.Itoa and printing its type.
//
// This file is intended for practice and reference on basic type conversions and error handling in Go.

func main() {

	var numericChar string = "12d"   // Declare a string variable with a non-numeric value
	var stringifiedBool string = "f" // Declare a string variable representing a boolean value
	asciiValue := 12

	parsedNumeric, err := strconv.Atoi(numericChar) // Try to convert numericChar to an integer
	// Declare an integer variable

	if err != nil { // Check if there was an error during conversion
		fmt.Println("Error parsing numericChar:", err) // Print error if conversion failed
	} else {
		fmt.Println("Parsed Numeric:", parsedNumeric) // Print the parsed integer if successful
	}

	fmt.Printf("Parsed Numeric: %d, Type of parsedNumeric: %T\n", parsedNumeric, parsedNumeric) // Print value and type

	// string to bool, parse boolean
	boolFlag, err := strconv.ParseBool(stringifiedBool) // Try to convert stringifiedBool to a boolean

	if err != nil { // Check if there was an error during conversion
		fmt.Println("Error parsing stringifiedBool:", err) // Print error if conversion failed
	}

	fmt.Println("Parsed Bool:", boolFlag) // Print the parsed boolean value

	// Integer to Ascii or string

	asciiString := strconv.Itoa(asciiValue)      // Convert integer to string
	fmt.Printf("ASCII Value: %T\n", asciiString) // Print the type of asciiString

}
