package main

import (
	"fmt"
)

func main() {

	// create map

	// method - 1
	colors1 := map[string]string {
		"red": "#ff0",
		"black": "#000",
		"green": "#111",
	}

	// method - 2 - not recommended
	var colors2 map[string]int

	// method - 3
	colors3 := make(map[string]string)
	colors3["orange"] = "#ffff"
	fmt.Print(colors1, colors2, colors3)

	// accessing key from map
	fmt.Println(colors1["red"])

	// delete key from map
	delete(colors1, "red")
	fmt.Println(colors1)


	printMap(colors1)

}


func printMap(c map[string]string) {
	for c, hex := range c {
		fmt.Println(c, hex)
	}
}