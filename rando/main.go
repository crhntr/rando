package main

import (
	"flag"

	"github.com/crhntr/rando"
)

func main() {
	var (
		length, number int
	)
	flag.IntVar(&length, "l", 3, "length of each section")
	flag.IntVar(&number, "n", 3, "number of sections")
	flag.Parse()

	println(rando.New(length, number))
}
