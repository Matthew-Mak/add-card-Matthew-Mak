package transaction

import (
	transactions "github.com/Matthew-Mak/card-Matthew-Mak/v2/pkg/types/transaction"
)

func SaveTransaction(CardId int, transaction transactions.Transaction, transactionsMap map[int][]transactions.Transaction) map[int][]transactions.Transaction {
	transactionsMap[CardId] = append(transactionsMap[CardId], transaction)
	return transactionsMap
}

func GetTransactionHistory(id int, transactionsMap map[int][]transactions.Transaction) []transactions.Transaction {
	if allTransactions, exists := transactionsMap[id]; exists {
		return allTransactions
	}
	return []transactions.Transaction{}
}
