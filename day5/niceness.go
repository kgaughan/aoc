package day5

import (
	"strings"
	"unicode/utf8"
)

func isVowel(ch rune) bool {
	return strings.ContainsRune("aeoiu", ch)
}

func IsNice(s string) bool {
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
