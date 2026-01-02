package main

import (
	"errors"
	"fmt"
)

func main() {

	str := "Go ^ Lang"

	for i, ch := range str {
		fmt.Printf("[character %d is %c] \n", i, ch)
	}

	str = appendString(str)

	fmt.Println("[post appending]", str)

	pointerAppend(&str)

	fmt.Println("[after pointer append]", str)

}

func pointerAppend(s *string) {
	if s == nil {
		errors.New("[empty string]")
	}

	*s = *s + "sujan"
}

func appendString(s string) string {

	s = s + " " + "complete guide"

	return s

}
