package withdraw

import (
	"errors"

	cards "github.com/Matthew-Mak/card-Matthew-Mak/v2/pkg/types/card"
)

var (
	ErrNegativeAmount   = errors.New("error: the amount can't be less than 0")
	ErrMoreThanExpected = errors.New("error: the amount can't higher than 100 000 000")
	ErrCardNotActive    = errors.New("error: card is not active")
	ErrNotEnoughBalance = errors.New("error: not enough money on the balance")
	MaximumAmount       = 100_000_000
)

func Withdraw(card *cards.Card, amount cards.Amount) (err error) {
	if amount < 0 {
		return ErrNegativeAmount
	}
	if amount > cards.Amount(MaximumAmount) {
		return ErrMoreThanExpected
	}
	if !card.IsActive {
		return ErrCardNotActive
	}
	if amount > cards.Amount(card.Balance) {
		return ErrNotEnoughBalance
	}
	card.Balance -= amount
	return nil

}
