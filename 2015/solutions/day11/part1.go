package day11

import "fmt"

func incrementString(runes []rune) []rune {
	offset := len(runes) - 1
	for {
		next := runes[offset] + 1
		if next > 'z' {
			runes[offset] = 'a'
			offset -= 1
			if offset == -1 {
				break
			}
		} else {
			runes[offset] = next
			break
		}
	}
	return runes
}

func checkRules(runes []rune, rules []func(runes []rune) bool) bool {
	for _, rule := range rules {
		if !rule(runes) {
			return false
		}
	}
	return true
}

func Part1(input string) {
	rules := []func(runes []rune) bool{
		hasAtLeastThreeAscending,
		hasNoConfusingCharacters,
		hasTwoPairs,
	}

	next := []rune(input)

	for {
		next = incrementString(next)
		if checkRules(next, rules) {
			break
		}
	}

	fmt.Printf("Next: %v\n", string(next))
}
