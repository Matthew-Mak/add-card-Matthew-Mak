package find

import (
	"errors"

	cards "github.com/Matthew-Mak/card-Matthew-Mak/v2/pkg/types/card"
)

var ErrCardIdNotFound = errors.New("error: the amount can't be less than 0")

func FindByID(id int, cardsSlice []cards.Card) (*cards.Card, error) {
	for i := 0; i < len(cardsSlice); i++ {
		if cardsSlice[i].Id == id {
			return &cardsSlice[i], nil
		}
	}
	return nil, ErrCardIdNotFound
}
