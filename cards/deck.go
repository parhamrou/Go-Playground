package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)
type deck []string


func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamons", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "One", "Two", "Three"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}
	return cards
}


func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}


func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}


func (d deck) toString() string{
	return strings.Join([]string(d), ",")
}


func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}


func readDeckFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}


func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	s := rand.New(source)
	for i := range d {
		newPosition := s.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}