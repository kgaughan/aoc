package day8

import (
	"fmt"
	"strconv"
)

func Part2(input string) {
	difference := parse(input, strconv.Quote)
	fmt.Printf("Difference is %v\n", difference)
}
