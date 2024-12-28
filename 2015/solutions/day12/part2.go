package day12

import (
	"fmt"
)

func Part2(input string) {
	fmt.Printf("Total is: %v\n", AddNumbers(parse(input), true))
}
