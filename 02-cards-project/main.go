package main

import "fmt"

func main() {
	// slice declaration
	cards := newDeck()
	// receiver function invocation
	cards.print()
	hand, remainingCards := deal(cards, 5)
	fmt.Println("[Hand Cards]")
	hand.print()
	fmt.Println("[Remaining Cards]")
	remainingCards.print()
	fmt.Println(cards.toString())
	fileName := "./data/cards"
	error := cards.saveFile(fileName)
	fmt.Println(error)
}
