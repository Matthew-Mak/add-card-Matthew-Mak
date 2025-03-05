package deposit_test

import (
	"fmt"

	"github.com/Matthew-Mak/add-card-Matthew-Mak/pkg/deposit"

	cards "github.com/Matthew-Mak/card-Matthew-Mak/v2/pkg/types/card"
)

func ExampleDeposit() {
	{
		// no errors test
		card := cards.Card{IsActive: true, Balance: 100}
		err := deposit.Deposit(&card, 100)
		fmt.Println(card, err)
	}

	{
		// card not active test
		card := cards.Card{IsActive: false, Balance: 100}
		err := deposit.Deposit(&card, 100)
		fmt.Println(card, err)
	}

	{
		// negative amount test
		card := cards.Card{IsActive: true, Balance: 100}
		err := deposit.Deposit(&card, -100)
		fmt.Println(card, err)
	}

	{
		// higher than 50 000 000 balance test test
		card := cards.Card{IsActive: true, Balance: 100}
		err := deposit.Deposit(&card, 50_000_000)
		fmt.Println(card, err)
	}
	{
		// higher than 50 000 000 balance test test
		card := cards.Card{IsActive: true, Balance: -100}
		err := deposit.Deposit(&card, 50)
		fmt.Println(card, err)
	}
	//Output: {0   200 true} <nil>
	//{0   100 false} error: card is not active
	//{0   100 true} error: the amount can't be less than 0
	//{0   100 true} error: the money on balance can't be more than 50 000 000
	//{0   -100 true} error: the balance after transfer is negative

}
