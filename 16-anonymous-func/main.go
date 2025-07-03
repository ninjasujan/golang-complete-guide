package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hello, World!")
	}()

	sum := find(1, 2, func(a, b int) int {
		return a + b
	})
	fmt.Println("sum:", sum)
}

func find(a, b int, f func(int, int) int) int {
	return f(a, b)
}
