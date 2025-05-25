package main

import "fmt"

func main() {

	var num1 uint8 = 10 // Declare an unsigned 8-bit integer variable num1 and assign value 10

	var ptr *uint8 = &num1 // Declare a pointer to uint8 and assign it the address of num1

	fmt.Println("pointer address:", ptr, "Pointer value:", *ptr) // Print the pointer's address and the value it points to
	fmt.Println("num1 value:", num1)                             // Print the value of num1

	num1 = 20 // Change the value of num1 to 20
	fmt.Println("After changing num1 value:")
	fmt.Println("pointer address:", ptr, "Pointer value:", *ptr) // Print pointer address and new value
	fmt.Println("num1 value:", num1)                             // Print updated num1 value

	Incr(ptr) // Call Incr function to increment the value pointed by ptr

	fmt.Println("After incrementing num1 using pointer:")
	fmt.Println("pointer address:", ptr, "Pointer value:", *ptr) // Print pointer address and incremented value

	// slice pointers

	list1 := make([]int, 5, 20)   // Create a slice of int with length 5 and capacity 20
	list2 := []int{1, 2, 3, 4, 5} // Create a slice of int with 5 elements

	list1[0] = 1 // Assign values to list1 elements
	list1[1] = 2
	list1[2] = 3
	list1[3] = 4
	list1[4] = 5

	fmt.Println(list1) // Print list1

	list1 = appendSlice(list1) // Call appendSlice to append elements and update list1

	fmt.Println(list1) // Print updated list1

	fmt.Println("list2 before append:", list2) // Print list2 before appending
	appendSliceViaPtr(&list2)                  // Call appendSliceViaPtr to append elements to list2 via pointer
	fmt.Println("list2 after append:", list2)  // Print list2 after appending

}

func Incr(num *uint8) {
	(*num)++ // Increment the value pointed by num
}

func appendSlice(slice []int) []int {

	slice = append(slice, 100, 200) // Append 100 and 200 to the slice

	slice[1] = 45 // Change the second element to 45

	return slice // Return the modified slice

}

func appendSliceViaPtr(slice *[]int) {
	*slice = append(*slice, 500, 600) // Append 500 and 600 to the slice via pointer
}
