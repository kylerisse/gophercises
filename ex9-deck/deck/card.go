package deck

import (
	"strings"
)

// Card in a deck has a suit and a rank
type Card struct {
	Suit rune
	Rank rune
}

// String for stringer interface
func (c Card) String() string {
	var sb strings.Builder
	sb.WriteRune('|')
	sb.WriteRune(c.Suit)
	sb.WriteRune(' ')
	if c.Rank == '0' {
		sb.WriteString("10|")
		return sb.String()
	}
	sb.WriteRune(' ')
	sb.WriteRune(c.Rank)
	sb.WriteRune('|')
	return sb.String()
}
