package main

import (
	"flag"
	"math/rand"
	"time"

	"github.com/crhntr/rando"
)

var (
	length, number int
)

func init() {
	rand.Seed(time.Now().UnixNano())

	flag.IntVar(&length, "l", 3, "length of each section")
	flag.IntVar(&number, "n", 3, "number of sections")
}

func main() {
	flag.Parse()

	println(rando.New(length, number))
}
