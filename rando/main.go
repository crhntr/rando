package main

import (
	"flag"

	"github.com/crhntr/rando"
)

var (
	length, number int
)

func init() {
	flag.IntVar(&length, "l", 3, "length of each section")
	flag.IntVar(&number, "n", 3, "number of sections")
}

func main() {
	flag.Parse()

	println(rando.New(length, number))
}
