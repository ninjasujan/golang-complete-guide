package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

func newDeckFromFile(fileName string) deck {
	deckByte, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("[Error]", err)
		os.Exit(1)
	}
	deckCards := strings.Split(string(deckByte), ",")
	return deck(deckCards)
}

func (cards deck) shuffle() {

	source := rand.NewSource(time.Now().Unix())
	rand := rand.New(source)

	for i := range cards {
		newPos := rand.Intn(len(cards) - 1)
		cards[i], cards[newPos] = cards[newPos], cards[i]
	}
}
