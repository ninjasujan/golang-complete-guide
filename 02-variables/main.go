package main

import (
	"fmt"
	"reflect"
)

// currentYear is a constant representing the current year.
// It's used to calculate the year of birth.
const currentYear = 2025 // Assuming current year is 2025 for the problem

// main is the entry point of the application.
// This function demonstrates:
// - Variable declarations (explicit type, type inference, grouped declaration)
// - Assigning and printing variable values
// - Simple calculations (year of birth)
// - Formatting output
// - Using reflection to print type information
// - Scoped variables (bonus example)
// - Basic type casting concepts
func main() {
	// Declare fullName using var with explicit type
	var fullName string = "Sujan"

	// Declare age using var with type inference (Go infers int type)
	var age = 30 // Change to your age

	// Grouped variable declaration for related variables
	var (
		isStudent      bool    = true // Whether the person is a student
		yearOfBirth    int            // Will be calculated below
		favoriteLetter string         // Defaults to empty string
		heightMeters   float64        // Defaults to 0.0
		gpa            float32        // Defaults to 0.0
	)

	// Calculate yearOfBirth using currentYear and age
	yearOfBirth = currentYear - age

	// Print the initial value of gpa (should be 0.0)
	fmt.Println("Initial GPA:", gpa)

	// Assign a value to gpa
	gpa = 3.7514441

	// Print the full profile using the variables above
	fmt.Println("Full Name:", fullName)
	fmt.Println("Age:", age)
	fmt.Println("Is Student:", isStudent)
	fmt.Println("Year of Birth:", yearOfBirth)
	fmt.Println("Favorite Letter:", favoriteLetter) // Will print empty string if not set
	fmt.Println("Height in Meters:", heightMeters)  // Will print 0 if not set
	fmt.Printf("GPA: %.2f\n", gpa)                  // Print GPA with 2 decimal places

	// --- Bonus Challenge Example ---
	// Demonstrates variable scoping: city is only accessible inside this block
	{
		city := "New York" // Change to your city
		fmt.Println("City:", city)
	}

	// Print the type of isStudent using reflection
	fmt.Printf("Type of isStudent: %s\n", reflect.TypeOf(isStudent).Kind())

	// Call printInfo function (demonstrates function with return value)
	printInfo()

	// Demonstrate rune and byte types (character types in Go)
	var char rune = 'A' // rune is an alias for int32, represents a Unicode code point
	var char2 byte = 65 // byte is an alias for uint8, represents an ASCII value

	fmt.Println("char:", string(char)) // Convert rune to string for display
	fmt.Println("char2:", char2)       // Prints numeric value (65)

	// Call typeCasting function (demonstrates type conversion concepts)
	typeCasting()
}

// printInfo demonstrates function definition, variable scope, and returning a pointer.
// It prints the sum of x and y, then returns a pointer to x.
func printInfo() *int {

	x := 10
	y := 20

	println("x:", x+y)

	return &x

}

// typeCasting demonstrates Go's strict type system and lack of implicit type casting.
// It also comments on the underlying types of rune and byte.
func typeCasting() {

	var a int = 10
	var b float32 = 10.76

	fmt.Println("a", a, "B", b)

	var results = float32(a) + b

	fmt.Println("Results:", results)

	// Note: No implicit type casting in Go.
	// rune is int32 and byte is uint8.
}
