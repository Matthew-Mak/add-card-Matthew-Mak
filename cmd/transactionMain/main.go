package main

import (
	"fmt"
	"time"

	cards "github.com/Matthew-Mak/card-Matthew-Mak/v2/pkg/types/card"
	transactions "github.com/Matthew-Mak/card-Matthew-Mak/v2/pkg/types/transaction"

	"github.com/Matthew-Mak/add-card-Matthew-Mak/pkg/display"
	"github.com/Matthew-Mak/add-card-Matthew-Mak/pkg/find"
	"github.com/Matthew-Mak/add-card-Matthew-Mak/pkg/transaction"
	"github.com/Matthew-Mak/add-card-Matthew-Mak/pkg/validate"
	"github.com/Matthew-Mak/add-card-Matthew-Mak/pkg/withdraw"
)

func main() {
	transactionsMap := make(map[int][]transactions.Transaction)
	var transactionElement transactions.Transaction
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
		fmt.Println("5: Show Transaction History.")
		fmt.Println("5: Exit.")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			panic(err)
		}

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
				_, err = fmt.Scan(&choiceAdd)
				if err != nil {
					panic(err)
				}
				switch choiceAdd {
				case "CARD":
					fmt.Print("Input card number: ")
					_, err = fmt.Scanln(&card.AccountID)
					if err != nil {
						panic(err)
					}

					fmt.Print("Input card date devided by (/): ")
					_, err = fmt.Scanln(&date)
					if err != nil {
						panic(err)
					}

					card, err = validate.Validate(card, date)
					if err != nil {
						panic(err)
					}
					fmt.Print("Input card balance: ")
					_, err = fmt.Scanln(&card.Balance)
					if err != nil {
						panic(err)
					}

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
			_, err := fmt.Scan(&inputAmount)
			if err != nil {
				panic(err)
			}

			fmt.Print("Input card Id: ")
			_, err = fmt.Scan(&id)
			if err != nil {
				panic(err)
			}

			cardFound, err := find.FindByID(id, cardSlice)
			if err != nil {
				panic(err)
			}
			err = withdraw.Withdraw(cardFound, cards.Amount(inputAmount))
			if err != nil {
				panic(err)
			}
			transactionElement = transactions.Transaction{Id: cardFound.Id, CardId: cardFound.AccountID, OperationType: "Withdraw", Amount: cards.Amount(inputAmount), Timestamp: time.Now()}
			transactionsMap = transaction.SaveTransaction(cardFound.Id, transactionElement, transactionsMap)
			fmt.Println("Resulted balance: ", &cardFound.Balance)
		case "4":
			// Пополнить баланс (Deposit)
			card := cards.Card{IsActive: true, Balance: 100}
			inputAmount := 0
			fmt.Println("Initial balance: ", card.Balance)
			fmt.Print("Input amount: ")
			_, err = fmt.Scan(&inputAmount)
			if err != nil {
				panic(err)
			}

			fmt.Print("Input card Id: ")
			_, err = fmt.Scan(&id)
			if err != nil {
				panic(err)
			}

			cardFound, err := find.FindByID(id, cardSlice)
			if err != nil {
				panic(err)
			}
			err = withdraw.Withdraw(cardFound, cards.Amount(inputAmount))
			if err != nil {
				panic(err)
			}
			transactionElement = transactions.Transaction{Id: cardFound.Id, CardId: cardFound.AccountID, OperationType: "Deposit", Amount: cards.Amount(inputAmount), Timestamp: time.Now()}
			transactionsMap = transaction.SaveTransaction(cardFound.Id, transactionElement, transactionsMap)
			fmt.Println("Resulted balance: ", &cardFound.Balance)
		case "5":
			fmt.Print("Input card Id: ")
			_, err = fmt.Scan(&id)
			if err != nil {
				panic(err)
			}

			transactinsForOutput := transaction.GetTransactionHistory(id, transactionsMap)
			for _, transact := range transactinsForOutput {
				fmt.Println(transactinsForOutput[transact.Id])
			}
		case "6":
			fmt.Println("Exiting program...")
			done = true
		}
	}
}
