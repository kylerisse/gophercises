package deck

import (
	"errors"
	"math/rand"
	"sort"
	"strings"
	"time"
)

// Deck of cards
type Deck struct {
	Cards  []Card
	refill []Card
}

const (
	// JOKER card
	JOKER = '🃏'
	// CLUBS suit
	CLUBS = '♣'
	// SPADES suit
	SPADES = '♠'
	// HEARTS suit
	HEARTS = '♥'
	// DIAMONDS suit
	DIAMONDS = '♦'
	// STARS suit
	STARS = '⭑'
)

var allCards = []Card{
	{'♣', '2'}, {'♣', '3'}, {'♣', '4'}, {'♣', '5'}, {'♣', '6'}, {'♣', '7'}, {'♣', '8'},
	{'♣', '9'}, {'♣', '0'}, {'♣', 'J'}, {'♣', 'Q'}, {'♣', 'K'}, {'♣', 'A'},
	{'♠', '2'}, {'♠', '3'}, {'♠', '4'}, {'♠', '5'}, {'♠', '6'}, {'♠', '7'}, {'♠', '8'},
	{'♠', '9'}, {'♠', '0'}, {'♠', 'J'}, {'♠', 'Q'}, {'♠', 'K'}, {'♠', 'A'},
	{'♥', '2'}, {'♥', '3'}, {'♥', '4'}, {'♥', '5'}, {'♥', '6'}, {'♥', '7'}, {'♥', '8'},
	{'♥', '9'}, {'♥', '0'}, {'♥', 'J'}, {'♥', 'Q'}, {'♥', 'K'}, {'♥', 'A'},
	{'♦', '2'}, {'♦', '3'}, {'♦', '4'}, {'♦', '5'}, {'♦', '6'}, {'♦', '7'}, {'♦', '8'},
	{'♦', '9'}, {'♦', '0'}, {'♦', 'J'}, {'♦', 'Q'}, {'♦', 'K'}, {'♦', 'A'},
	{'⭑', '2'}, {'⭑', '3'}, {'⭑', '4'}, {'⭑', '5'}, {'⭑', '6'}, {'⭑', '7'}, {'⭑', '8'},
	{'⭑', '9'}, {'⭑', '0'}, {'⭑', 'J'}, {'⭑', 'Q'}, {'⭑', 'K'}, {'⭑', 'A'},
	{'🃏', '*'}, {'🃏', '*'},
}

// NewDeck creates a new deck with filter
func NewDeck(filter func([]Card) []Card) *Deck {
	cc := filter(allCards)
	return &Deck{
		Cards:  cc,
		refill: cc,
	}
}

// Draw a card from the deck, get error if empty
func (d *Deck) Draw() (Card, error) {
	if len(d.Cards) < 1 {
		return Card{}, errors.New("Deck is empty")
	}
	c := d.Cards[0]
	d.Cards = d.Cards[1:]
	return c, nil
}

// Shuffle collect all cards and shuffle the deck
func (d *Deck) Shuffle() {
	cc := d.refill
	rand.Seed(time.Now().UTC().UnixNano())
	for i := range cc {
		j := rand.Intn(i + 1)
		cc[i], cc[j] = cc[j], cc[i]
	}
	d.Cards = cc
}

// Sort the deck
func (d *Deck) Sort() {
	sort.Sort(BySuit(d.Cards))
}

// String for stringer interface
func (d *Deck) String() string {
	var sb strings.Builder
	for i, c := range d.Cards {
		if i%5 == 0 && i != 0 {
			sb.WriteString("\n")
		}
		sb.WriteString(c.String() + " ")
	}
	sb.WriteString("\n")
	return sb.String()
}
