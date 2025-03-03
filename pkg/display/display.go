package display

import (
	"fmt"

	"github.com/Matthew-Mak/add-card-Matthew-Mak/pkg/validate"
)

func DisplaySliceOfCards(cards []validate.Card) {
	if len(cards) == 0 {
		fmt.Println("No cards in the system...")
	} else {
		for i := 0; i < len(cards); i++ {
			fmt.Printf("Card: id: %v, gate: %v, pan: %v - Balance: %v\n", cards[i].Id, cards[i].Number, cards[i].System, cards[i].Balance)
		}
	}
}
