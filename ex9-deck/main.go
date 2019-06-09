package main

import (
	"fmt"

	"./deck"
)

func main() {
	deck1 := deck.NewDeck(deck.BlackJack(6))
	deck1.Sort()
	fmt.Println(deck1)
	fmt.Println("--------")
	deck1.Shuffle()
	fmt.Println(deck1)
	fmt.Println("--------")
	deck2 := deck.NewDeck(deck.FiveCrowns)
	fmt.Println(deck2)
	fmt.Println("------")
	deck2.Shuffle()
	fmt.Println(deck2)
}
