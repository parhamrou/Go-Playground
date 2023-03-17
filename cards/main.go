package main

//import "fmt"

//import "fmt"

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}

func newCard() string { 
	return "Five of Diomonds"
}