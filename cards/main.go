package main

import "fmt"

func main() {
	// var card string = "Ace of spades" // long form of initializing a var
	cards := deck{newCard(), newCard()}    // slice of cards
	cards = append(cards, "Six of spades") // returns new slice

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	cards.print()

	fmt.Println(cards)
}

func newCard() string {
	return "Five of diamonds"
}
