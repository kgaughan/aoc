package main

import (
	"flag"
	"fmt"
	"math"

	"aoc/day4"
)

func main() {
	n := flag.Int("n", 5, "How long should the string of zeroes be?")
	flag.Parse()

	prefix := flag.Arg(0)
	if prefix == "" {
		fmt.Println("You must provide a prefix")
		return
	}

	answer := day4.AdventCoin(prefix, *n, 1, math.MaxInt32)
	fmt.Printf("%q for %v is %v\n", prefix, *n, answer)
}
