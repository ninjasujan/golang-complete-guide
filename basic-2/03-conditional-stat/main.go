package main

import "fmt"


func main() {

	a := 10
	b := 20

	if a < b {
		fmt.Println(a)
	} else {
		fmt.Println(b)
	}

	// || && ! - logical operator

	if a >= 10 && a < 20 {
		fmt.Println("A passed the test")
	}

	if a <= 10 {
		fmt.Println("a")
	} else if a > 100 {
		fmt.Println("b")
	}


}




