package main

import "fmt"

func main() {

	var num1 uint8 = 10

	var ptr *uint8 = &num1

	fmt.Println("pointer address:", ptr, "Pointer value:", *ptr)
	fmt.Println("num1 value:", num1)

	num1 = 20
	fmt.Println("After changing num1 value:")
	fmt.Println("pointer address:", ptr, "Pointer value:", *ptr)
	fmt.Println("num1 value:", num1)

	Incr(ptr)

	fmt.Println("After incrementing num1 using pointer:")
	fmt.Println("pointer address:", ptr, "Pointer value:", *ptr)

}

func Incr(num *uint8) {
	(*num)++
}
