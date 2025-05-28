package util

import (
	"fmt"
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

func RandomFullName() string {
	return fmt.Sprintf("%s %s", RandomString(6), RandomString(6))
}

func RandomEmail() string {
	return fmt.Sprintf("%s@%s.com", RandomString(8), RandomString(6))
}
