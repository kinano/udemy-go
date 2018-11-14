package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades and got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected blah but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deck_testing")
	d := newDeck()
	d.saveToFile("_deck_testing")
	loadedD, _ := newDeckFromFile("_deck_testing")
	if len(loadedD) != 16 {
		t.Errorf("Expecting 16 got %v", len(loadedD))
	}
}
