package main

import (
	"fmt"
	"reflect"
)

const currentYear = 2025 // Assuming current year is 2025 for the problem

func main() {
	// Declare fullName using var with explicit type
	var fullName string = "Sujan"

	// Declare age using var with type inference
	var age = 30 // Change to your age

	// TODO: Declare other variables (isStudent, yearOfBirth, favoriteLetter, heightMeters, gpa)

	var (
		isStudent      bool = true
		yearOfBirth    int
		favoriteLetter string
		heightMeters   float64
		gpa            float32
	)

	// TODO: Calculate yearOfBirth

	yearOfBirth = currentYear - age

	// TODO: Print initial gpa

	fmt.Println("Initial GPA:", gpa)

	// TODO: Assign a value to gpa

	gpa = 3.7514441

	// TODO: Print the full profile
	fmt.Println("Full Name:", fullName)
	fmt.Println("Age:", age)
	fmt.Println("Is Student:", isStudent)
	fmt.Println("Year of Birth:", yearOfBirth)
	fmt.Println("Favorite Letter:", favoriteLetter)
	fmt.Println("Height in Meters:", heightMeters)
	fmt.Printf("GPA: %.2f\n", gpa)

	// --- Bonus Challenge Example ---
	{
		city := "New York" // Change to your city
		fmt.Println("City:", city)
	}

	fmt.Printf("Type of fullName: %s\n", reflect.TypeOf(isStudent).Kind())

	printInfo()
}

func printInfo() *int {

	x := 10
	y := 20

	println("x:", x+y)

	return &x

}
