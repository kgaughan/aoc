package day5

import (
	"bufio"
	"strings"
)

func parse(input string, niceness func(string) bool) int {
	r := strings.NewReader(input)
	scanner := bufio.NewScanner(r)
	nice := 0
	for scanner.Scan() {
		if niceness(scanner.Text()) {
			nice++
		}
	}
	return nice
}
