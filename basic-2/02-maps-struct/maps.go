package main

import "fmt"

func main () {
	// method -1
	statePopulation := map[string]int {
		"Texas": 45123,
		"Florida": 42333,
		"New York": 3122,
		"Ohio": 412222,
	}

	// method -2
	countryPopulation := make(map[string]int)
	countryPopulation = map[string]int {
		"Texas": 45123,
		"Florida": 42333,
		"New York": 3122,
		"Ohio": 412222,
	}

	// operation on map - delete key
	delete(statePopulation, "Texas")

	// length of map
	fmt.Println(len(countryPopulation))



	fmt.Println(statePopulation, countryPopulation)
}

