package validate

import (
	"errors"
	cards "github.com/Matthew-Mak/card-Matthew-Mak/v2/pkg/types/card"
	"strconv"
	"strings"
	"time"
)

var (
	ErrNumberTooLong     = errors.New("error: the card number is longer than 16 symbols")
	ErrNumberTooShort    = errors.New("error: the card number is shorter than 16 symbols")
	ErrNumberHasNoDigits = errors.New("error: the input must consist only of digits")
	ErrWrongCardSystem   = errors.New("error: card system is invalid")
	ErrCardYearExpired   = errors.New("error: card date year is expired")
	ErrCardMonthExpired  = errors.New("error: card date month is expired")
)

type CardDate string

func Validate(card cards.Card, date CardDate) (cards.Card, error) {
	var (
		currentMonth time.Month
		currentYear  int
	)
	currentYear, currentMonth, _ = time.Now().Date()

	const (
		uzcardSystemNumber = "8600"
		humoSystemNumber   = "9860"
		uzcardSystemName   = "UzCard"
		humoSystemName     = "Humo"
	)

	if len(card.AccountID) > 16 {
		return card, ErrNumberTooLong
	}
	if len(card.AccountID) < 16 {
		return card, ErrNumberTooShort
	}
	for i := 0; i < len(card.AccountID); i++ {
		if int(card.AccountID[i]) > 57 || int(card.AccountID[i]) < 48 {
			return card, ErrNumberHasNoDigits
		}
	}

	//date validation
	convertedDate := strings.Split(string(date), "/")
	cardYear, errCardYear := strconv.Atoi(convertedDate[0])
	if errCardYear != nil {
		return card, errCardYear
	}
	cardMonth, errCardMonth := strconv.Atoi(convertedDate[1])
	if errCardMonth != nil {
		return card, errCardMonth
	}

	if cardYear < currentYear-((currentYear/100)*100) { //Converting year, e.g. 2025 into 25
		return card, ErrCardYearExpired
	} else if cardMonth < int(currentMonth) {
		return card, ErrCardMonthExpired
	}

	if card.AccountID[:4] == uzcardSystemNumber {
		card.Pan = uzcardSystemName
	} else if card.AccountID[:4] == humoSystemNumber {
		card.Pan = humoSystemName
	} else {
		return card, ErrWrongCardSystem
	}

	card.Id = 1
	//Replacing numbers with xes
	card.AccountID = card.AccountID[:4] + "xxxxxxxx" + card.AccountID[12:]
	//Inserting spaces
	card.AccountID = card.AccountID[:4] + " " + card.AccountID[4:8] + " " + card.AccountID[8:12] + " " + card.AccountID[12:]

	return card, nil
}
