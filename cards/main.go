package main

import "fmt"

func square(x *float64) {
	*x = *x * *x
}

func main() {
	cards := newDeck()

	hand, remainingDeck := deal(cards, 5)

	// cards.print()
	hand.print()
	remainingDeck.print()

	cards2 := newDeck()
	cards2.saveToFile("my_cards")

	cards3 := newDeckFromFile("my_cards")
	cards3.print()

	cards4 := newDeck()
	cards4.shuffle()
	cards4.print()

	i := 10

	if i > 10 {
		fmt.Println("Big")
	} else {
		fmt.Println("Small")
	}

	for i := 0; i < 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}

	for i := 0; i < 100; i++ {
		switch {
		case (i%3 == 0) && (i%5 == 0):
			fmt.Println(i, "FizzBuzz")
		case i%3 == 0:
			fmt.Println(i, "Fizz")
		case i%5 == 0:
			fmt.Println(i, "Buzz")
		default:
			fmt.Println(i)
		}
	}

	x := 1.5
	fmt.Println((&x))
}
