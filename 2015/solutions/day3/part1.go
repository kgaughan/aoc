package day3

import (
	"fmt"
)

func Part1(input string) {
	visited := deliver(input, 1)
	fmt.Printf("Houses visited: %v\n", visited)
}
