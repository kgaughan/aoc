package day5

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func Part1(input string) {
	fmt.Printf("nice entries: %v\n", parse(input, isNice1))
}

func isNice1(s string) bool {
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
