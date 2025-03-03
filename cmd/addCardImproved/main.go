package main

import (
	"fmt"

	"github.com/Matthew-Mak/add-card-Matthew-Mak/pkg/display"
	"github.com/Matthew-Mak/add-card-Matthew-Mak/pkg/validate"
)

func main() {
	cards := make([]validate.Card, 0, 5)
	var card validate.Card
	var date validate.CardDate
	done := false
	choice := ""

	for !done {

		// Добавить карту
		if len(cards) == 5 {
			done = true
			fmt.Println("Card limit is reached, exiting program...")
		}
		fmt.Print("Input CARD to add your card or EXIT to quit: ")
		fmt.Scanln(&choice)
		switch choice {
		case "CARD":
			fmt.Print("Input card number: ")
			fmt.Scanln(&card.Number)
			fmt.Print("Input card date devided by (/): ")
			fmt.Scanln(&date)
			card, err := validate.Validate(card, date)
			if err != nil {
				panic(err)
			}
			fmt.Print("Input card balance: ")
			fmt.Scanln(&card.Balance)
			cards = append(cards, card)
		case "EXIT":
			done = true
			fmt.Println("Exiting program...")
		}
	}

	// Показать карты
	display.DisplaySliceOfCards(cards)

	//// Снять деньги (Withdraw)
	// card := types.Card{IsActive: true, Balance: 100}
	// input := 0
	// fmt.Println("Initial balance: ", card.Balance)
	// fmt.Print("Input amount: ")
	// fmt.Scan(&input)
	// err := withdraw.Withdraw(&card, types.Amount(input))
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Resulted balance: ", card.Balance)

	// Пополнить баланс (Deposit)

}
