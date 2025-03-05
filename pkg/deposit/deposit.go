package deposit

import (
	"errors"

	cards "github.com/Matthew-Mak/card-Matthew-Mak/v2/pkg/types/card"
)

var (
	ErrNegativeAmount   = errors.New("error: the amount can't be less than 0")
	ErrCardNotActive    = errors.New("error: card is not active")
	ErrNegativeBalance  = errors.New("error: the balance after transfer is negative")
	ErrTooMuchOnBalance = errors.New("error: the money on balance can't be more than 50 000 000")
	MaximumBalance      = 50_000_000
)

func Deposit(card *cards.Card, amount cards.Amount) (err error) {
	if amount < 0 {
		return ErrNegativeAmount
	}
	if !card.IsActive {
		return ErrCardNotActive
	}
	if amount+cards.Amount(card.Balance) < 0 {
		return ErrNegativeBalance
	}
	if amount+cards.Amount(card.Balance) > cards.Amount(MaximumBalance) {
		return ErrTooMuchOnBalance
	}
	card.Balance += amount
	return nil

}
