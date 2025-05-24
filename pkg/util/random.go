package util

import (
	"math/rand/v2"
)

func RandomInt(min, max int64) int64 {
	return rand.Int64N(max-min+1) + min
}

func RandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyz"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.IntN(len(letters))]
	}
	return string(b)
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	return currencies[rand.IntN(len(currencies))]
}
