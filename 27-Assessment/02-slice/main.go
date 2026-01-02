package main

import "fmt"

func main() {

	marks := make([]int, 6)

	populateMarks(marks)

	fmt.Println("[student mark]", marks)

	for i, mark := range marks {
		fmt.Println("[student marks]", i, mark)
	}

}

func populateMarks(m []int) {

	m[0] = 99
	m[1] = 98
	m[2] = 76
	m[3] = 89
	m[4] = 67
	m[5] = 67

}
