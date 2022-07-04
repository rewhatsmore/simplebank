package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandString(n int) string {
	var sb strings.Builder
	l := len(alphabet)
	for i := 0; i < n; i++ {
		s := alphabet[rand.Intn(l)]
		sb.WriteByte(s)
	}
	return sb.String()
}

// RandomOwner generates a random owner name
func RandOwner() string {
	return RandString(6)
}

// RandomMoney generates a random amount of money
func RandMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generates a random currency code
func RandCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	return currencies[rand.Intn(len(currencies))]
}
