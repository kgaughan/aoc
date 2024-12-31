package day8

import (
	"strings"

	"github.com/kgaughan/aoc/2015/helpers"
)

func parse(input string, conv func(original string) string) int {
	originalLength := 0
	convertedLength := 0
	helpers.ScanLines(strings.NewReader(input), func(original string) error {
		originalLength += len(original)
		convertedLength += len(conv(original))
		return nil
	})

	difference := originalLength - convertedLength
	if originalLength < convertedLength {
		difference = -difference
	}

	return difference
}
