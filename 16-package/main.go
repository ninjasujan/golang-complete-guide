package main

import (
	"fmt"
	"shapes/rect"
)

func main() {

	newRect := rect.NewRect(10, 10)

	fmt.Println("[Area of a Rectangle]: ", newRect.Area())

	fmt.Println("[Perimeter of a Rectangle]:", newRect.Perimeter())

}
