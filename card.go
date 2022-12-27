package card

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func Type(cardNumber string) string {
	prefix, _ := strconv.Atoi(cardNumber[:2])

	switch prefix {
		case 34 , 37:
			return "American Express"
		case 51 , 52 , 53 , 54 , 55:
			return "MasterCard"
		case 40, 41, 42, 43, 44, 45, 46, 47, 48, 49:
			return "Visa"
	}
	return "unknown"
}

func Length(cardNumber string) int {
	return utf8.RuneCountInString(strings.Replace(cardNumber, " ", "", -1))
}

func luhnAlgorithm(cardNumber string) bool {
	var firstSum int
	var secondSum int

	for i := Length(cardNumber) -2; i >= 0; i -= 2 {
		n, _ := strconv.Atoi(string(cardNumber[i]))
		nd := n * 2

		if nd > 9 {
			firstSum += nd / 10
			firstSum += nd % 10
		} else {
			firstSum += nd
		}
	}

	for i := 1; i < Length(cardNumber); i += 2 {
		n, _ := strconv.Atoi(string(cardNumber[i]))
		secondSum += n
	}

	return (firstSum + secondSum) % 10 == 0
}

func IsValid(cardNumber string) bool {
	return luhnAlgorithm(strings.Replace(cardNumber, " ", "", -1))
}

func Details(cardNumber string) {
	var information strings.Builder

	information.WriteString(fmt.Sprintf("Card Number:\t%s\n",	cardNumber))
	information.WriteString(fmt.Sprintf("Card Type:\t%s\n", 	Type(cardNumber)))
	information.WriteString(fmt.Sprintf("Card Length:\t%d\n",	Length(cardNumber)))
	information.WriteString(fmt.Sprintf("Card Valid:\t%v",		IsValid(cardNumber)))

	fmt.Println(information.String())
}
