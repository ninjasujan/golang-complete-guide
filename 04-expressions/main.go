package main

import (
	"fmt"
	"math"
)

func main() {

	a, b, c := 12, 7, 3

	expResult1 := math.Pow(float64(a), 2) + math.Pow(float64(b), 2) + 2*float64(a)*float64(b)
	expResult2 := a*b + c/a - b%c

	fmt.Println("Expression Result 1: ", expResult1)
	fmt.Println("Expression Result 2: ", expResult2)

	fmt.Println("Is a > b: ", a > b)
	fmt.Println("Is c equal to a '%' b", c == (a%b))

	// sqrt of c

	fmt.Println("Square root of c: ", math.Sqrt(float64(c)))

}
