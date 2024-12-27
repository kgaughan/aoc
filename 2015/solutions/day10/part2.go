package day10

import (
	"fmt"
)

func Part2(input string) {
	convoluted := input
	for i := 0; i < 50; i++ {
		convoluted = LookAndSay(convoluted)
	}
	fmt.Printf("%v\n", len(convoluted))
}
