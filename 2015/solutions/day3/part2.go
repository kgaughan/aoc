package day3

import (
	"fmt"
)

func Part2(input string) {
	visited := deliver(input, 2)
	fmt.Printf("Houses visited: %v\n", visited)
}
