package day8

import (
	"fmt"
	"strconv"
)

func Part1(input string) {
	difference := parse(input, func(s string) string {
		unquoted, _ := strconv.Unquote(s)
		return unquoted
	})
	fmt.Printf("Difference is %v\n", difference)
}
