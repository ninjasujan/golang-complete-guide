package main

import "fmt"

func main() {

	var name = "sujan poojary 你好"

	// Print each byte as a character.
	// This works for ASCII, but not for Unicode characters (like 你 or 好),
	// because Unicode characters can take more than one byte.
	for i := 0; i < len(name); i++ {
		fmt.Println("[char] -> ", string(name[i]))
	}

	// Use a for-range loop to iterate over the string.
	// This extracts each Unicode code point (rune) correctly,
	// so multi-byte characters are handled properly.
	for index, char := range name {
		fmt.Println("[idx] -> ", index, " [Char] -> ", string(char))
	}

	// Convert the string to a slice of runes.
	// This allows direct indexing of Unicode characters,
	// so each element in the rune slice is a full character.
	runeString := []rune(name)
	for i := 0; i < len(runeString); i++ {
		fmt.Println("[char] -> ", string(runeString[i]))
	}

}
