package validate

import (
	"errors"
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

type Card struct {
	Id      int
	Number  string
	System  string
	Balance int
}

func Validate(card Card, date CardDate) (Card, error) {
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

	if len(card.Number) > 16 {
		return card, ErrNumberTooLong
	}
	if len(card.Number) < 16 {
		return card, ErrNumberTooShort
	}
	for i := 0; i < len(card.Number); i++ {
		if int(card.Number[i]) > 57 || int(card.Number[i]) < 48 {
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

	if card.Number[:4] == uzcardSystemNumber {
		card.System = uzcardSystemName
	} else if card.Number[:4] == humoSystemNumber {
		card.System = humoSystemName
	} else {
		return card, ErrWrongCardSystem
	}

	card.Id = 1
	//Replacing numbers with xes
	card.Number = card.Number[:4] + "xxxxxxxx" + card.Number[12:]
	//Inserting spaces
	card.Number = card.Number[:4] + " " + card.Number[4:8] + " " + card.Number[8:12] + " " + card.Number[12:]

	return card, nil
}
