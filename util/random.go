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

// RandomInt generates a random int between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomOwner returns a random owner name
func RandomOwner() string {
	return RandomString(10)
}

// RandomCash returns a random amount of money
func RandomCash() int64 {
	return RandomInt(0, 10000)
}

// RandomCurrency returns a random currency from predefined list
func RandomCurrency() string {
	curr := []string{"USD", "CAD", "ZAR", "ZWE", "JPY", "SEK"}
	return curr[rand.Intn(len(curr))]
}
