package display_test

import (
	"github.com/Matthew-Mak/add-card-Matthew-Mak/pkg/display"
	cards "github.com/Matthew-Mak/card-Matthew-Mak/v2/pkg/types/card"
)

func ExampleDisplaySliceOfCards() {
	// no errors test
	{
		card1 := cards.Card{Id: 1, AccountID: "8600123412341234", Pan: "UzCard", Balance: 500, IsActive: true}
		card2 := cards.Card{Id: 2, AccountID: "8600123412341234", Pan: "UzCard", Balance: 500, IsActive: true}
		card3 := cards.Card{Id: 3, AccountID: "8600123412341234", Pan: "UzCard", Balance: 500, IsActive: true}

		cards := make([]cards.Card, 0, 5)
		cards = append(cards, card1)
		cards = append(cards, card2)
		cards = append(cards, card3)

		display.DisplaySliceOfCards(cards)
	}
	//Output: Card: id: 1, gate: 8600123412341234, pan: UzCard - Balance: 500
	//Card: id: 2, gate: 8600123412341234, pan: UzCard - Balance: 500
	//Card: id: 3, gate: 8600123412341234, pan: UzCard - Balance: 500
}
