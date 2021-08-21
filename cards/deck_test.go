package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 36 {
		t.Errorf("Expected deck length of 36, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of spades, but got %v", d[0])
	}

	if d[len(d) - 1] != "Nine of Clubs" {
		t.Errorf("Expected last card of Nine of Clubs, but got %v", d[0])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 36 {
		t.Errorf("Expected deck length of 36, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}