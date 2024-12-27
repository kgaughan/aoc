package day5

import (
	"strings"
)

func isVowel(ch rune) bool {
	return strings.ContainsRune("aeoiu", ch)
}

// Check for pairs of identical characters separated by another character.
func hasTwins(rs []rune) bool {
	for i := 0; i < len(rs)-2; i++ {
		if rs[i] == rs[i+2] {
			return true
		}
	}
	return false
}
