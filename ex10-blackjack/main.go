package main

import (
	"flag"
	"fmt"
	"os"

	"./blackjack"
)

func main() {
	var numDecks int
	flag.IntVar(&numDecks, "d", 6, "number of decks")
	flag.Parse()
	if numDecks < 1 {
		fmt.Printf("Invalid number of decks %d\n", numDecks)
		os.Exit(1)
	}
	game := blackjack.NewCli(numDecks)
	game.Run()
}
