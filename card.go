package card

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type card struct {
	number string
}

func New(number string) card {
	return card{
		number: number,
	}
}

func (c card) Details() {
	var information strings.Builder

	information.WriteString(fmt.Sprintf("Card Number:\t%s\n", c.number))
	information.WriteString(fmt.Sprintf("Card Type:\t%s\n", c.Network()))
	information.WriteString(fmt.Sprintf("Card Length:\t%d\n", c.length()))
	information.WriteString(fmt.Sprintf("Card Valid:\t%v", c.IsValid()))

	fmt.Println(information.String())
}

func (c card) IsValid() bool {
	return c.luhnAlgorithm()
}

func (c card) Network() string {
	prefix, _ := strconv.Atoi(c.number[:2])

	switch prefix {
	case 34, 37:
		return "American Express"
	case 51, 52, 53, 54, 55:
		return "MasterCard"
	case 40, 41, 42, 43, 44, 45, 46, 47, 48, 49:
		return "Visa"
	}
	return "unknown"
}

func (c card) length() int {
	return utf8.RuneCountInString(strings.Replace(c.number, " ", "", -1))
}

func (c card) luhnAlgorithm() bool {
	var firstSum int
	var secondSum int

	cc := strings.Replace(c.number, " ", "", -1)

	for i := c.length() - 2; i >= 0; i -= 2 {
		n, _ := strconv.Atoi(string(cc[i]))
		nd := n * 2

		if nd > 9 {
			firstSum += nd / 10
			firstSum += nd % 10
		} else {
			firstSum += nd
		}
	}

	for i := 1; i < c.length(); i += 2 {
		n, _ := strconv.Atoi(string(cc[i]))
		secondSum += n
	}

	return (firstSum+secondSum)%10 == 0
}
