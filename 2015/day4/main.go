package main

import (
	"flag"
	"fmt"
	"math"

	"github.com/kgaughan/aoc/2015/day4/lib"
)

func main() {
	n := flag.Int("n", 5, "How long should the string of zeroes be?")
	flag.Parse()

	prefix := flag.Arg(0)
	if prefix == "" {
		fmt.Println("You must provide a prefix")
		return
	}

	answer := lib.AdventCoin(prefix, *n, 1, math.MaxInt32)
	fmt.Printf("%q for %v is %v\n", prefix, *n, answer)
}
