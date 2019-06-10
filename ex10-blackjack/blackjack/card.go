package blackjack

import (
	"strings"
)

// Card in a deck has a suit and a rank
type card struct {
	suit rune
	rank rune
}

// String for stringer interface
func (c card) String() string {
	var sb strings.Builder
	sb.WriteRune('|')
	sb.WriteRune(c.suit)
	sb.WriteRune(' ')
	if c.rank == '0' {
		sb.WriteString("10|")
		return sb.String()
	}
	sb.WriteRune(' ')
	sb.WriteRune(c.rank)
	sb.WriteRune('|')
	return sb.String()
}
