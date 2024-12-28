package day16

import "fmt"

func Part1(input string) {
	sues := parse(input)
	matched := findSue(sues, func(_ string, value, expected int) bool {
		return expected == value
	})
	fmt.Printf("%v\n", matched)
}
