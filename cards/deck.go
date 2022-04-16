package main

import "fmt"

// create a new type of deck
// which is a slice of strings
type deck []string // deck extends all functionality from []string

// define receiver
// any variable of type "deck"
// now gets access to the
// "print" method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
