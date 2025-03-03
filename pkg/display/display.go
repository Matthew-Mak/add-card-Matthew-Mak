package display

import (
	"fmt"
	cards "github.com/Matthew-Mak/card-Matthew-Mak/v2/pkg/types/card"
)

func DisplaySliceOfCards(cards []cards.Card) {
	if len(cards) == 0 {
		fmt.Println("No cards in the system...")
	} else {
		for i := 0; i < len(cards); i++ {
			fmt.Printf("Card: id: %v, gate: %v, pan: %v - Balance: %v\n", cards[i].Id, cards[i].AccountID, cards[i].Pan, cards[i].Balance)
		}
	}
}
