package day10

import (
	"fmt"
)

func Part1(input string) {
	convoluted := input
	for i := 0; i < 40; i++ {
		convoluted = LookAndSay(convoluted)
	}
	fmt.Printf("%v\n", len(convoluted))
}
