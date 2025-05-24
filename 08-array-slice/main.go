package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// Array creation
	var list = [5]int{1, 2, 3, 4, 5}
	var name string = "AAAA"
	var sliceList1 []int
	var sliceList2 = make([]int, 5)
	var sliceList3 = []int{1, 2, 3}

	// Both strings are deep copied since string is immutable.
	var name2 = name
	// Print string header address, data pointer, and length for name and name2
	fmt.Printf("String 1 header: %p, data ptr: %p, len: %d\n", &name, unsafe.Pointer((*[2]uintptr)(unsafe.Pointer(&name))[0]), len(name))
	fmt.Printf("String 2 header: %p, data ptr: %p, len: %d\n", &name2, unsafe.Pointer((*[2]uintptr)(unsafe.Pointer(&name2))[0]), len(name2))

	name2 = "BBBB" // assign a new value to name2 (creates a new string, different data pointer)
	fmt.Printf("String 2 header: %p, data ptr: %p, len: %d\n", &name2, unsafe.Pointer((*[2]uintptr)(unsafe.Pointer(&name2))[0]), len(name2))

	// Iterate over array and print values
	for _, val := range list {
		fmt.Println("[value of array at index]:", val)
	}

	// String operation
	strOp(name)
	fmt.Println("[str:]", name)
	fmt.Println("[str alias:]", name2)

	// Slice data structure
	fmt.Println("[size of a slice]", unsafe.Sizeof(sliceList1))
	fmt.Println("[Slice 2]", sliceList2)
	fmt.Println("[Slice 3]", sliceList3)
	fmt.Printf("slice header: %p ptr: %p, len: %d, cap: %d\n", &sliceList3, &sliceList3[0], len(sliceList3), cap(sliceList3))

	// Slice operations
	testSlice1 := []int{1, 2, 3, 4, 5}
	testSlice2 := testSlice1

	fmt.Printf("[testSlice1 before modification] %p: %v\n", testSlice1, testSlice1)
	fmt.Printf("[testSlice2 before modification] %p: %v\n", testSlice2, testSlice2)

	testSlice1 = append(testSlice2, 6)

	fmt.Printf("[testSlice1 after modification] %p: %v\n", testSlice1, testSlice1)
	fmt.Printf("[testSlice2 after modification] %p: %v\n", testSlice2, testSlice2)

	sliceTest := []int{}               // initialize an empty slice without specifying length
	sliceTest = append(sliceTest, 100) // append value to the slice
	fmt.Println("[sliceTest] ", sliceTest[0])

	fmt.Println("[sliceTest] ", sliceTest[0])
}

// strOp demonstrates string immutability by modifying a copy of the string
func strOp(s string) {
	s += "padding...."                   // append to string (creates new string, does not affect original)
	fmt.Println("[str inside strop]", s) // print modified string
}
