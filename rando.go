package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	N, S int
)

func init() {
	rand.Seed(time.Now().UnixNano())

	flag.IntVar(&N, "n", 3, "enter the number of chars pre section")
	flag.IntVar(&S, "s", 3, "enter the number of sections")
}

func main() {
	flag.Parse()

	for i := 0; i < S; i++ {
		fmt.Print(randGen(N))
		if i < S-1 {
			fmt.Print("-")
		}
	}
	fmt.Println()
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randGen(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
