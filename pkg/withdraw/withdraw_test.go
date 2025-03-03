package withdraw_test

import (
	"fmt"

	"github.com/Matthew-Mak/add-card-Matthew-Mak/pkg/withdraw"
	cards "github.com/Matthew-Mak/card-Matthew-Mak/v2/pkg/types/card"
)

func ExampleWithdraw() {
	{
		// no errors test
		card := cards.Card{IsActive: true, Balance: 200}
		err := withdraw.Withdraw(&card, 100)
		fmt.Println(card, err)
	}

	{
		//card not active error
		card := cards.Card{IsActive: false, Balance: 200}
		err := withdraw.Withdraw(&card, 100)
		fmt.Println(card, err)

	}

	{
		// negative amount test
		card := cards.Card{IsActive: true, Balance: 100}
		err := withdraw.Withdraw(&card, -100)
		fmt.Println(card, err)
	}

	{
		// amount error
		card := cards.Card{IsActive: true, Balance: 100}
		err := withdraw.Withdraw(&card, 101_000_000)
		fmt.Println(card, err)
	}

	{
		//not enough balance error
		card := cards.Card{IsActive: true, Balance: 50}
		err := withdraw.Withdraw(&card, 100)
		fmt.Println(card, err)
	}

	{
		// no errors test with 0 balance
		card := cards.Card{IsActive: true, Balance: 100}
		err := withdraw.Withdraw(&card, 100)
		fmt.Println(card, err)
	}

	//Output: {0   100 true} <nil>
	//{0   200 false} error: card is not active
	//{0   100 true} error: the amount can't be less than 0
	//{0   100 true} error: the amount can't higher than 100 000 000
	//{0   50 true} error: not enough money on the balance
	//{0   0 true} <nil>
}
