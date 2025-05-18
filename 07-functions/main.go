package main

import (
	"errors"
	"fmt"
)

func main() {

	var input = "sujan" // Define an integer input

	// Call isOddOrEven with 8 (not input), check if it's even, and handle errors
	isEven, err := isOddOrEven(input)

	if err != nil {
		fmt.Println("[Input error]", err) // Print error if input is invalid
	} else {
		// Print result: is the given number even?
		fmt.Println("Given number: ", input, " is even?", isEven)
	}

}

// isOddOrEven checks if the given number is even.
// Returns true if even, false if odd, and error if input is not an integer.
func isOddOrEven(num any) (bool, error) {

	// Check if the type of num is int using reflection
	intVal, ok := num.(int)
	if !ok {
		return false, errors.New("not a valid number type")
	}

	// Return true if num is even, false otherwise
	return intVal%2 == 0, nil

}
