package day12

import (
	"fmt"
)

func Part1(input string) {
	fmt.Printf("Total is: %v\n", AddNumbers(parse(input), false))
}
