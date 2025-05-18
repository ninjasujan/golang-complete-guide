package main

import (
	"fmt"
	"unsafe"
)

func main() {

	// array creation
	var list = [5]int{1, 2, 3, 4, 5} // create an array of 5 integers
	var name string = "AAAA"         // create a string
	var sliceList1 []int             // declare a nil slice
	var sliceList2 = make([]int, 5)  // create a slice of length 5 (all zeroes)
	var sliceList3 = []int{1, 2, 3}  // create a slice with 3 elements

	/**
		both strings are deep copied since string is immutable assigning any value to string creates new string and points to string header, so old one will be discarded
	**/

	var name2 = name // copy string to name2 (both point to same data initially)
	// print string header address, data pointer, and length for name and name2
	fmt.Printf("String 1 header: %p, data ptr: %p, len: %d\n", &name, unsafe.Pointer((*[2]uintptr)(unsafe.Pointer(&name))[0]), len(name))
	fmt.Printf("String 2 header: %p, data ptr: %p, len: %d\n", &name2, unsafe.Pointer((*[2]uintptr)(unsafe.Pointer(&name2))[0]), len(name2))

	name2 = "BBBB" // assign a new value to name2 (creates a new string, different data pointer)

	// print header and data pointer for name2 after reassignment
	fmt.Printf("String 2 header: %p, data ptr: %p, len: %d\n", &name2, unsafe.Pointer((*[2]uintptr)(unsafe.Pointer(&name2))[0]), len(name2))

	// array initializations (redundant, already initialized above)
	list[0] = 1
	list[1] = 2
	list[2] = 3
	list[3] = 4
	list[4] = 5
	// list[5] = 6 // out of 	bound

	// iterate over array and print values
	for _, val := range list {
		fmt.Println("[value of array at index]: ", val)
	}

	/** string operation **/
	/**
		when string is reassigned or modified, a new string is created,
	 with same header but internal data pointer and length will be changed
	**/
	strOp(name)                        // call strOp with name (does not modify original string)
	fmt.Println("[str:]", name)        // print original string (unchanged)
	fmt.Println("[str alias:]", name2) // print name2 (may be changed)

	/** Slice data structure
		slice Header
			Pointer
			Len
			Capacity
		Slice is pass by value in Go, its header is passed by value
	**/
	fmt.Println("[size of a slice]", unsafe.Sizeof(sliceList1)) // print size of slice header (24 bytes on 64-bit)
	fmt.Println("[Slice 2]", sliceList2)                        // print sliceList2 (all zeroes)
	fmt.Println("[Slice 3]", sliceList3)                        // print sliceList3 ([1 2 3])

	// print slice header address, pointer to first element, length, and capacity
	fmt.Printf("slice header: %p ptr: %p, len: %d, cap: %d\n", &sliceList3, &sliceList3[0], len(sliceList3), cap(sliceList3))

}

// strOp demonstrates string immutability by modifying a copy of the string
func strOp(s string) {
	s += "padding...."                   // append to string (creates new string, does not affect original)
	fmt.Println("[str inside strop]", s) // print modified string
}
