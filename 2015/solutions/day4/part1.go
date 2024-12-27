package day4

import (
	"flag"
	"fmt"
	"math"
)

var zeroes = flag.Int("day4zeroes", 5, "How long should the string of zeroes be?")

func Part1(input string) {
	answer := adventCoin(input, *zeroes, 1, math.MaxInt32)
	fmt.Printf("%q for %v is %v\n", input, *zeroes, answer)
}
