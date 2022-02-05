package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// create new type of deck - slice of string
type deck []string

func newDeck() deck {
	cards := deck{}
	cardsSuits := []string{"Spades", "Dimonds", "Hearts", "Clubs"}
	cardsValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardsSuits {
		for _, value := range cardsValues {
			cards = append(cards, suit+" of "+value)
		}
	}
	return cards
}

// function of type deck prints the each card of the cards - receiver function
func (cards deck) print() {
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

// function multiple return value
func deal(cards deck, handSize int) (deck, deck) {
	return cards[:handSize], cards[handSize:]
}

func (cards deck) toString() string {
	return strings.Join([]string(cards), ",")
}

func (cards deck) saveFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(cards.toString()), 0644)
}
