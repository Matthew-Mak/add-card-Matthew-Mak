package validate_test

import (
	"fmt"
	"main/pkg/validate"
)

func ExampleValidate() {
	// no errors test
	{
		card := validate.Card{1, "8600123412341234", "", 500}
		var date validate.CardDate = "26/11"
		card, err := validate.Validate(card, date)
		fmt.Println(card, err)
	}
	// expired year error test
	{
		card := validate.Card{1, "8600123412341234", "", 500}
		var date validate.CardDate = "22/11"
		card, err := validate.Validate(card, date)
		fmt.Println(card, err)
	}
	// expired month error test
	{
		card := validate.Card{1, "8600123412341234", "", 500}
		var date validate.CardDate = "25/1"
		card, err := validate.Validate(card, date)
		fmt.Println(card, err)
	}
	// invalid system error test
	{
		card := validate.Card{1, "8500123412341234", "", 500}
		var date validate.CardDate = "26/11"
		card, err := validate.Validate(card, date)
		fmt.Println(card, err)
	}
	//Output: {1 8600 xxxx xxxx 1234 UzCard 500} <nil>
	//{1 8600123412341234  500} error: card date year is expired
	//{1 8600123412341234  500} error: card date month is expired
	//{1 8500123412341234  500} error: card system is invalid
}
