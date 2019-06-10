package blackjack

type game struct {
	dealer   hand
	player   hand
	shoe     shoe
	wins     int
	losses   int
	pushes   int
	numDecks int
}

func newGame(numDecks int) *game {
	return &game{
		dealer:   *newHand(),
		player:   *newHand(),
		numDecks: numDecks,
		shoe: shoe{
			shuffleTime: true,
		},
	}
}

func (g *game) nextHand() {
	g.player = *newHand()
	g.dealer = *newHand()
	g.player.add(g.shoe.draw())
	g.dealer.add(g.shoe.draw())
	g.player.add(g.shoe.draw())
	g.dealer.add(g.shoe.draw())
}

func (g *game) shuffle() {
	g.shoe = *newShoe(g.numDecks)
}
