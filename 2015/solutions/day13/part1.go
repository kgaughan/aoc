package day13

import (
	"fmt"
	"strings"
)

func Part1(input string) {
	entries, err := parse(strings.NewReader(input))
	if err != nil {
		panic(err)
	}
	best := tryAll(entries, false)
	fmt.Printf("Best: %v\n", best)
}
