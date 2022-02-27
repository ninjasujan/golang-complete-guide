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

	// file IO operation
	fileName := "./data/cards.txt"
	error := cards.saveFile(fileName)
	fmt.Println(error)

	newDeckFromFile(fileName).print()

	// shuffle cards
	cards.shuffle()
	cards.print()
}
