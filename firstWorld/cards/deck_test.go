package main

import (
	"os"
	"testing"
)

// to create a test in go we have to create a file called "{xxxnamexxx}_test.go"
// later, $: go test -> for run all the testsg

// tests functions must start with a upper-case word named: Test+funcName
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Cloud" {
		t.Errorf("Expected last card of Fourd of Cloud but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 chars in deck but got %v", len(loadedDeck))
	}

	// a file with _ it's telling us that is a temporary file
	os.Remove("_decktesting")
}
