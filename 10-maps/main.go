package main

import (
	"fmt"
	"strings"
)

func main() {
	// Print a header for the program
	fmt.Println("[Maps in Go]")

	// Create a map to store string keys and int values
	mapList := make(map[string]int)

	// Add a key-value pair to the map
	mapList["one"] = 1

	// Print the current state of the map
	fmt.Println("mapList:", mapList)

	// Retrieve the value for the key "one" and check if it exists
	val, ok := mapList["one"]
	if ok {
		fmt.Println("[Key Found in the map]", val)
	}

	// Define a string to count character occurrences
	str := "my Name is sujan"

	// Create a map to count occurrences of each character
	mapCounter := make(map[rune]int) // Use rune for better Unicode support

	// Iterate over each character in the string
	for _, char := range str {
		char = rune(strings.ToLower(string(char))[0])
		mapCounter[char]++ // Increment the count for this character
	}

	// Print the character count map
	fmt.Println("[count of char in the string]")

	// Print each character and its count for better readability
	for char, count := range mapCounter {
		fmt.Printf("'%c': %d\n", char, count)
	}
}
