package main

import "fmt"

func main () {
	names := [3]string {"john", "david"} // array
	grades := []int {67, 56, 67}
	var matrix [3][3]int = [3][3]int { {1, 0, 1},  {1, 0, 1},  {1, 0, 1}}
	var test  = make([]int, 3)
	// grades2 := [...]int {98, 56, 78} please asssign size to hold the data which has passed
	fmt.Println(names, matrix, grades, test)
	

	/**

		- len(array) - function to get length of the array
		- copying arary makes deep copy in golang
		- To make shallow copy use &in source array
		- append(srcArray, ...eleToAdd)
		- 
	**/
}

