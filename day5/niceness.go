package day5

import (
	"strings"
	"unicode/utf8"
)

func isVowel(ch rune) bool {
	return strings.ContainsRune("aeoiu", ch)
}

type NiceFunc func(string) bool

func IsNice1(s string) bool {
	naughties := []string{
		"ab",
		"cd",
		"pq",
		"xy",
	}
	for _, naughty := range naughties {
		if strings.Contains(s, naughty) {
			return false
		}
	}

	nVowels := 0

	lastCh, offset := utf8.DecodeRuneInString(s)
	if isVowel(lastCh) {
		nVowels++
	}

	hasNiceSeq := false
	for _, ch := range s[offset:] {
		if isVowel(ch) {
			nVowels++
		}
		if ch == lastCh {
			hasNiceSeq = true
		}
		lastCh = ch
	}

	return nVowels >= 3 && hasNiceSeq
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

// Check for pairs of identical characters separated by another character.
func hasTwins(rs []rune) bool {
	for i := 0; i < len(rs)-2; i++ {
		if rs[i] == rs[i+2] {
			return true
		}
	}
	return false
}

func IsNice2(s string) bool {
	rs := []rune(s)
	return len(rs) > 3 && hasDupes(rs) && hasTwins(rs)
}
