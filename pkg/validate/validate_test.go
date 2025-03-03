package validate_test

import (
	"fmt"
	cards "github.com/Matthew-Mak/card-Matthew-Mak/v2/pkg/types/card"

	"github.com/Matthew-Mak/add-card-Matthew-Mak/pkg/validate"
)

func ExampleValidate() {
	// no errors test UzCard
	{
		card := cards.Card{Id: 1, AccountID: "8600123412341234", Pan: "", Balance: 500, IsActive: true}
		var date validate.CardDate = "26/11"
		card, err := validate.Validate(card, date)
		fmt.Println(card, err)
	}
	// no errors test Humo
	{
		card := cards.Card{Id: 1, AccountID: "9860123412341234", Pan: "", Balance: 500, IsActive: true}
		var date validate.CardDate = "26/11"
		card, err := validate.Validate(card, date)
		fmt.Println(card, err)
	}
	// expired year error test
	{
		card := cards.Card{Id: 1, AccountID: "8600123412341234", Pan: "", Balance: 500, IsActive: true}
		var date validate.CardDate = "22/11"
		card, err := validate.Validate(card, date)
		fmt.Println(card, err)
	}
	// expired month error test
	{
		card := cards.Card{Id: 1, AccountID: "8600123412341234", Pan: "", Balance: 500, IsActive: true}
		var date validate.CardDate = "25/1"
		card, err := validate.Validate(card, date)
		fmt.Println(card, err)
	}
	// invalid system error test
	{
		card := cards.Card{Id: 1, AccountID: "8500123412341234", Pan: "", Balance: 500, IsActive: true}
		var date validate.CardDate = "26/11"
		card, err := validate.Validate(card, date)
		fmt.Println(card, err)
	}
	//Output: {1 8600 xxxx xxxx 1234 UzCard 500 true} <nil>
	//{1 9860 xxxx xxxx 1234 Humo 500 true} <nil>
	//{1 8600123412341234  500 true} error: card date year is expired
	//{1 8600123412341234  500 true} error: card date month is expired
	//{1 8500123412341234  500 true} error: card system is invalid
}
