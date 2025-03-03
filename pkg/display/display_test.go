package display_test

import (
	"github.com/Matthew-Mak/add-card-Matthew-Mak/pkg/display"
	"github.com/Matthew-Mak/add-card-Matthew-Mak/pkg/validate"
)

func ExampleDisplaySliceOfCards() {
	// no errors test
	{
		card1 := validate.Card{Id: 1, Number: "8600123412341234", System: "UzCard", Balance: 500}
		card2 := validate.Card{Id: 2, Number: "8600123412341234", System: "UzCard", Balance: 500}
		card3 := validate.Card{Id: 3, Number: "8600123412341234", System: "UzCard", Balance: 500}

		cards := make([]validate.Card, 0, 5)
		cards = append(cards, card1)
		cards = append(cards, card2)
		cards = append(cards, card3)

		display.DisplaySliceOfCards(cards)
	}
	//Output: Card: id: 1, gate: 8600123412341234, pan: UzCard - Balance: 500
	//Card: id: 2, gate: 8600123412341234, pan: UzCard - Balance: 500
	//Card: id: 3, gate: 8600123412341234, pan: UzCard - Balance: 500
}
