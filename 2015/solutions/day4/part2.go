package day4

import (
	"fmt"
	"math"
)

func Part2(input string) {
	answer := adventCoin(input, 6, 1, math.MaxInt32)
	fmt.Printf("%q is %v\n", input, answer)
}
