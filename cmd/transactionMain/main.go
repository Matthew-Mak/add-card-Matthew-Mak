package main

import (
	"fmt"
	"github.com/Matthew-Mak/add-card-Matthew-Mak/pkg/display"
	"github.com/Matthew-Mak/add-card-Matthew-Mak/pkg/find"
	"github.com/Matthew-Mak/add-card-Matthew-Mak/pkg/validate"
	"github.com/Matthew-Mak/add-card-Matthew-Mak/pkg/withdraw"
	cards "github.com/Matthew-Mak/card-Matthew-Mak/v2/pkg/types/card"
)
import transactions "github.com/Matthew-Mak/card-Matthew-Mak/v2/pkg/types/transaction"

func main() {
	transactions := make(map[int][]transactions.Transaction)
	cardSlice := make([]cards.Card, 0, 5)
	var card cards.Card
	var date validate.CardDate
	done := false
	choice := ""
	id := 0

	for !done {
		fmt.Println("Please input numbers 1-5 to choose operation:")
		fmt.Println("1: Add card.")
		fmt.Println("2: Show cards.")
		fmt.Println("3: Withdraw card.")
		fmt.Println("4: Deposit card.")
		fmt.Println("5: Exit.")
		fmt.Scanln(&choice)

		switch choice {

		case "1":
			doneAdd := false
			choiceAdd := ""

			for !doneAdd {
				// Добавить карту
				if len(cardSlice) == 5 {
					doneAdd = true
					fmt.Println("Card limit is reached, exiting program...")
				}
				fmt.Print("Input CARD to add your card or EXIT to quit: ")
				fmt.Scanln(&choiceAdd)
				switch choiceAdd {
				case "CARD":
					fmt.Print("Input card number: ")
					fmt.Scanln(&card.AccountID)
					fmt.Print("Input card date devided by (/): ")
					fmt.Scanln(&date)
					card, err := validate.Validate(card, date)
					if err != nil {
						panic(err)
					}
					fmt.Print("Input card balance: ")
					fmt.Scanln(&card.Balance)
					cardSlice = append(cardSlice, card)
				case "EXIT":
					doneAdd = true
					fmt.Println("Exiting program...")
				}
			}
		case "2":
			// Показать карты
			display.DisplaySliceOfCards(cardSlice)
		case "3":
			// Снять деньги (Withdraw)
			inputAmount := 0
			fmt.Println("Initial balance: ", card.Balance)
			fmt.Print("Input amount: ")
			fmt.Scan(&inputAmount)
			fmt.Print("Input card Id: ")
			fmt.Scan(&id)

			cardFound, err := find.FindByID(id, cardSlice)
			if err != nil {
				panic(err)
			}
			err = withdraw.Withdraw(cardFound, cards.Amount(inputAmount))
			if err != nil {
				panic(err)
			}
			fmt.Println("Resulted balance: ", &cardFound.Balance)
		case "4":
			// Пополнить баланс (Deposit)
			card := cards.Card{IsActive: true, Balance: 100}
			inputAmount := 0
			fmt.Println("Initial balance: ", card.Balance)
			fmt.Print("Input amount: ")
			fmt.Scan(&inputAmount)
			fmt.Print("Input card Id: ")
			fmt.Scan(&id)

			cardFound, err := find.FindByID(id, cardSlice)
			if err != nil {
				panic(err)
			}
			err = withdraw.Withdraw(cardFound, cards.Amount(inputAmount))
			if err != nil {
				panic(err)
			}
			fmt.Println("Resulted balance: ", &cardFound.Balance)
		case "5":
			fmt.Println("Exiting program...")
			done = true
		}
	}
}
