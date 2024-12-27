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

// Check if the string contains a sequence of two characters two or more
// times without overlap.
func hasDupes(rs []rune) bool {
	for i := 0; i < len(rs)-2; i++ {
		for j := i + 2; j < len(rs)-1; j++ {
			if rs[i] == rs[j] && rs[i+1] == rs[j+1] {
				return true
			}
		}
	}
	return false
}
