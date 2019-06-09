package deck

import "sort"

// Standard52 is a standard deck of 52 cards with no jokers
var Standard52 = func(cc []Card) []Card {
	newCards := []Card{}
	for _, c := range cc {
		if c.Suit == '⭑' || c.Suit == '🃏' {
			continue
		}
		newCards = append(newCards, c)
	}
	return newCards
}

// BlackJack returns a multi-deck deck
var BlackJack = func(num int) func(cc []Card) []Card {
	return func(cc []Card) []Card {
		newCards := []Card{}
		for num > 0 {
			for _, c := range Standard52(allCards) {
				newCards = append(newCards, c)
			}
			num--
		}
		sort.Sort(BySuit(newCards))
		return newCards
	}
}

// Standard54 is a standard deck of 52 cards with jokers
var Standard54 = func(cc []Card) []Card {
	newCards := []Card{}
	for _, c := range cc {
		if c.Suit == '⭑' {
			continue
		}
		newCards = append(newCards, c)
	}
	return newCards
}

// FiveCrowns is a Five Crowns deck
var FiveCrowns = func(cc []Card) []Card {
	newCards := []Card{}
	for _, c := range cc {
		if c.Rank == '2' {
			continue
		}
		newCards = append(newCards, c)
	}
	return newCards
}
