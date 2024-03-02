package util

import (
	"math/rand"
)

const (
	Prefix    = "HBD"
	MinAmount = 10_000
	MaxAmount = 1_000_000
)

var numbers = []rune("1234567890")

// source => https://stackoverflow.com/a/22892986
func GeneratePromoCode(n int) string {
	code := make([]rune, n)

	for i := range code {
		code[i] = numbers[rand.Intn(len(numbers))]
	}

	return Prefix + string(code)
}

// source => https://stackoverflow.com/a/23577092
func GeneratePromoAmount() int {
	return rand.Intn(MaxAmount-MinAmount) + MinAmount
}
