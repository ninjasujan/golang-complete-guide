package main

import "fmt"

// interfaces
type bot interface {
	getGreetings() string
}

// struct
type englishBot struct {}
type spanishBot struct {}

func main() {
	engBot := englishBot {}
	spanBot := spanishBot {}
	printGreetings(engBot)
	printGreetings(spanBot)
}

func printGreetings(b bot) {
	fmt.Println(b.getGreetings())
}

func (eBot englishBot) getGreetings() string {
	return "Hello There.."
}

func (sBot spanishBot) getGreetings() string {
	return "Hawla.!"
}