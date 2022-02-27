package main

import "testing"

func TestNewDecj(t *testing.T) {
	deck := newDeck()

	if len(deck) != 16 {
		t.Errorf("Expected deck length of 16 but %d", len(deck))
	}

	if deck[0] != "Spades of Ace" {
		t.Errorf("Expected first card of Ace of spade, recevied %v", deck[0])
	}

	if deck[len(deck)-1] != "Clubs of Four" {
		t.Errorf("Expected first card of Ace of spade, recevied %v", deck[0])
	}

}
