package main

import "fmt"

func main() {

	fmt.Println("[Define Custom Type and Methods]")

	var numType MyInt = 10

	fmt.Println(numType)

	numType.print()

}

type MyInt int

func (num MyInt) print() {
	fmt.Println(num)
}
