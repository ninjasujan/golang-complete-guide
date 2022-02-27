package main

import "fmt"


func main () {


	/**

		Implicitly breaks were applied in golang switch
		if you want fallthrough behaviour in switch blocks use
		fallthrough keyword at the end

	**/

	key := 2

	switch key {
	case 1: 
		fmt.Println(key)
		// fallthrough
	case 2:
		fmt.Println(key)
	default:
		fmt.Println(key, )
	}


}