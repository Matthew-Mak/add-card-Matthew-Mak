package find_test

import (
	"fmt"
	"github.com/Matthew-Mak/add-card-Matthew-Mak/pkg/find"
	cards "github.com/Matthew-Mak/card-Matthew-Mak/v2/pkg/types/card"
)

func ExampleFindByID() {
	// no errors test
	{
		card1 := cards.Card{Id: 1, AccountID: "8600123412341234", Pan: "UzCard", Balance: 500, IsActive: true}
		card2 := cards.Card{Id: 2, AccountID: "8600123412341234", Pan: "UzCard", Balance: 500, IsActive: true}
		card3 := cards.Card{Id: 3, AccountID: "8600123412341234", Pan: "UzCard", Balance: 500, IsActive: true}

		cards := make([]cards.Card, 0, 5)
		cards = append(cards, card1)
		cards = append(cards, card2)
		cards = append(cards, card3)

		foundCard, err := find.FindByID(1, cards)
		fmt.Println(&foundCard, err)
	}
	// invalid Id test
	{
		card1 := cards.Card{Id: 1, AccountID: "8600123412341234", Pan: "UzCard", Balance: 500, IsActive: true}
		card2 := cards.Card{Id: 2, AccountID: "8600123412341234", Pan: "UzCard", Balance: 500, IsActive: true}
		card3 := cards.Card{Id: 3, AccountID: "8600123412341234", Pan: "UzCard", Balance: 500, IsActive: true}

		cards := make([]cards.Card, 0, 5)
		cards = append(cards, card1)
		cards = append(cards, card2)
		cards = append(cards, card3)

		foundCard, err := find.FindByID(4, cards)
		fmt.Println(foundCard, err)
	}
	//Output: 0xc00004c0d0 <nil>
	//<nil> error: the amount can't be less than 0
}
