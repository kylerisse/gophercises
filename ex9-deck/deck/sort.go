package deck

var suitValues = map[rune]int{
	'♣': 0, '♠': 20, '♥': 40, '♦': 60, '⭑': 80, '🃏': 100,
}
var rankValues = map[rune]int{
	'2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8,
	'9': 9, '0': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14,
}

// BySuit is a sort interface for cards
type BySuit []Card

// Len of []card
func (b BySuit) Len() int {
	return len(b)
}

// Less of []card
func (b BySuit) Less(i, j int) bool {
	return suitValues[b[i].Suit]+rankValues[b[i].Rank] <
		suitValues[b[j].Suit]+rankValues[b[j].Rank]
}

// Swap of []card
func (b BySuit) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
