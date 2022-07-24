package utils

import (
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var macAddressRunes = []rune("ABCDEF0123456789")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func RandMacAddress() string {
	b := make([]rune, 17)
	for i := range b {
		if i > 0 && i < 17 && i%3 == 0 {
			b[i] = ':'
		} else {
			b[i] = macAddressRunes[rand.Intn(len(macAddressRunes))]
		}
	}
	return string(b)
}
