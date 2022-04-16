package main

import "fmt"

// create a new type of deck
// which is a slice of strings
type deck []string // deck extends all functionality from []string

func newDeck() deck {
	cards := deck{}

	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "J", "Q", "K"}

	for _, cardSuite := range cardSuites {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+" of "+cardSuite)
		}
	}

	return cards
}

// define receiver
// any variable of type "deck"
// now gets access to the
// "print" method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
