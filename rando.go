package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	N int
)

func init() {
	rand.Seed(time.Now().UnixNano())

	flag.IntVar(&N, "n", 3, "enter the number of chars pre section")
}

func main() {
	flag.Parse()
	fmt.Printf("%s-%s-%s\n", randGen(N), randGen(N), randGen(N))
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ012345678901234567890123456789")

func randGen(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
