package main

import "testing"
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected len to be 16, but it's %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but it's %v", d[0])
	}
}


func TestWorkWithFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile()
	loadedDeck := readDeckFile()
	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards, but it's %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}