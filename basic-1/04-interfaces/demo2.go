package main

import "fmt"

func main() {
	count := IntCounter(0)
	inc := &count
	for i := 0; i < 10; i++  {
		fmt.Println(inc.Increment())
	}
}


type Incrementer interface {
	Increment() int
}

type IntCounter int


func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}	