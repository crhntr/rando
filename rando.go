package rando

import (
	"math/rand"
	"strings"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randGen(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// New creates a random string built of sections separated by a '-'
func New(length, number int) string {
	var str strings.Builder
	for i := 0; i < number; i++ {
		str.WriteString(randGen(length))
		if i < number-1 {
			str.WriteString("-")
		}
	}
	return str.String()
}
