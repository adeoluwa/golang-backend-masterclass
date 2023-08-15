package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// The RandomInt function generates a random integer between a given minimum and maximum value.
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max - min + 1) // min ->max
}

// The function RandomString generates a random string of length n using characters from the alphabet.
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++{
		c:= alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string{
	return RandomString(6)
}

//RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

//RandomCurrency generates a random currency
func RandomCurrency() string {
	currencies := []string{EUR, USD, CAD, NGN}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}


// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@gmail.com", RandomString(6))
}