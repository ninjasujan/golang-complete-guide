package main

import (
	"fmt"
)

func main() {
	const myConst int = 56 // can't initialize const to something which determines at run time
	var val float32 = 4
	fmt.Printf("%v, %T\n", myConst, myConst)
	sum := val + float32(myConst)
	fmt.Printf("%v\n", sum)
}


