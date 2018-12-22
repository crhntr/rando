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

// New creates a random string with letters and numbers
func New(sectionLen, sectionsCount int) string {
	var str strings.Builder
	for i := 0; i < sectionsCount; i++ {
		str.WriteString(randGen(sectionLen))
		if i < sectionsCount-1 {
			str.WriteString("-")
		}
	}
	return str.String()
}
