package day13

import (
	"fmt"
	"io"
	"strings"
)

func parse(source io.Reader) (map[pair]int, error) {
	entries := make(map[pair]int, 100)
	for {
		if pair, happiness, err := parseEntry(source); err == nil {
			entries[pair] = happiness
		} else if err == io.EOF {
			return entries, nil
		} else {
			return nil, err
		}
	}
}

func parseEntry(source io.Reader) (pair, int, error) {
	var change string
	var person string
	var other string
	var happiness int
	_, err := fmt.Fscanf(
		source,
		"%s would %s %d happiness units by sitting next to %s\n",
		&person, &change, &happiness, &other)
	if err == nil && change == "lose" {
		happiness *= -1
	}
	return pair{person: person, other: strings.TrimRight(other, ".")}, happiness, err
}
