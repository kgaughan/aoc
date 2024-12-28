package day13

import (
	"fmt"
	"strings"
)

func Part2(input string) {
	entries, err := parse(strings.NewReader(input))
	if err != nil {
		panic(err)
	}
	best := tryAll(entries, true)
	fmt.Printf("Best: %v\n", best)
}
