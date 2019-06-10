package blackjack

import (
	"strconv"
	"strings"
)

type hand struct {
	cards []card
}

func newHand() *hand {
	return &hand{}
}

func (h *hand) add(c card) {
	h.cards = append(h.cards, c)
}

func (h *hand) points() int {
	var aces int
	var points int
	for _, c := range h.cards {
		switch c.rank {
		case 'K', 'Q', 'J', '0':
			points += 10
		case 'A':
			aces++
		default:
			val, err := strconv.Atoi(string(c.rank))
			if err != nil {
				panic(err)
			}
			points += val
		}
	}
	for aces > 1 {
		points++
		aces--
	}
	if aces > 0 && points < 11 {
		points += 11
		aces--
	}
	if aces > 0 {
		points++
		aces--
	}
	return points
}

// String for stringer interface
func (h *hand) String() string {
	var sb strings.Builder
	for _, c := range h.cards {
		sb.WriteString(c.String() + " ")
	}
	return sb.String()
}
