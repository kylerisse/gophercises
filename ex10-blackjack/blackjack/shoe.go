package blackjack

import (
	"math/rand"
	"strings"
	"time"
)

type shoe struct {
	cards       []card
	shuffleTime bool
}

func newShoe(numDecks int) *shoe {
	shoeCards := []card{}
	// add multiple decks per parameter
	for ; numDecks > 0; numDecks-- {
		shoeCards = append(shoeCards, fullDeck...)
	}
	// shuffle
	rand.Seed(time.Now().UTC().UnixNano())
	for i := range shoeCards {
		j := rand.Intn(i + 1)
		shoeCards[i], shoeCards[j] = shoeCards[j], shoeCards[i]
	}
	// insert the cut card somewhere between 50% and 70%
	numShoeCards := len(shoeCards)
	cutCardSpot := rand.Intn(numShoeCards*2/10) + (numShoeCards * 5 / 10)
	second := append([]card{cutCard}, shoeCards[cutCardSpot:]...)
	first := shoeCards[0:cutCardSpot]
	shoeCards = append(first, second...)
	// return the shoe to caller
	return &shoe{
		cards: shoeCards,
	}
}

func (s *shoe) draw() card {
	c := s.cards[0]
	if c == cutCard {
		s.shuffleTime = true
		s.cards = s.cards[1:]
		c = s.cards[0]
	}
	s.cards = s.cards[1:]
	return c
}

// String for stringer interface
func (s *shoe) String() string {
	var sb strings.Builder
	for i, c := range s.cards {
		if i%5 == 0 && i != 0 {
			sb.WriteString("\n")
		}
		sb.WriteString(c.String() + " ")
	}
	sb.WriteString("\n")
	return sb.String()
}
