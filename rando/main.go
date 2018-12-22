package main

import (
	"flag"
	"math/rand"
	"time"

	"github.com/crhntr/rando"
)

var (
	sectionLen, sectionsCount int
)

func init() {
	rand.Seed(time.Now().UnixNano())

	flag.IntVar(&sectionLen, "sl", 3, "section-len is the number of chars pre section")
	flag.IntVar(&sectionsCount, "sc", 3, "section-count is enter the number of sections")
}

func main() {
	flag.Parse()

	println(rando.New(sectionLen, sectionsCount))
}
