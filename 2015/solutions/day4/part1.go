package day4

import (
	"fmt"
	"math"
)

func Part1(input string) {
	answer := adventCoin(input, 1, 1, math.MaxInt32)
	fmt.Printf("%q is %v\n", input, answer)
}
