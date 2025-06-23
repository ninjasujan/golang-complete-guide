package main

import "fmt"

func main() {

	fmt.Println("[Go Generics -  Generic Functions]")

	var intList []int
	var stringList []string

	intList = appendElement(intList, 4)
	stringList = appendElement(stringList, "Hello")

	fmt.Println(intList)
	fmt.Println(stringList)

	fmt.Println("[Go Generics  - Generic Slice]")

	var container Container[int] = &Stack[int]{}
	container.Append(4)
	container.Append(5)
	container.Append(6)
	container.Remove()

	fmt.Println(container)

}

type Stack[T any] []T

type Container[T any] interface {
	Append(T)
	Remove()
	Len() int
}

func appendElement[T any](list []T, ele T) []T {
	return append(list, ele)
}

func (st *Stack[T]) Append(ele T) {
	*st = append(*st, ele)
}

func (st *Stack[T]) Remove() {
	if len(*st) == 0 {
		return
	}
	*st = (*st)[:len(*st)-1]
}

func (st *Stack[T]) Len() int {
	return len(*st)
}
