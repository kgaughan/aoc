package day5

import "fmt"

func Part2(input string) {
	fmt.Printf("nice entries: %v\n", parse(input, isNice2))
}

func isNice2(s string) bool {
	rs := []rune(s)
	return len(rs) > 3 && hasDupes(rs) && hasTwins(rs)
}
