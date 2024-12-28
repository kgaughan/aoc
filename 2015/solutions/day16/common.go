package day16

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

var expectedProperties = map[string]int{
	"children":    3,
	"cats":        7,
	"samoyeds":    2,
	"pomeranians": 3,
	"akitas":      0,
	"vizslas":     0,
	"goldfish":    5,
	"trees":       3,
	"cars":        2,
	"perfumes":    1,
}

func findSue(sues map[int]map[string]int, fn func(key string, value, expected int) bool) int {
	for sue, properties := range sues {
		good := true
		for key, expected := range expectedProperties {
			if value, ok := properties[key]; ok && !fn(key, value, expected) {
				good = false
			}
		}
		if good {
			return sue
		}
	}
	return 0
}

func parseEntry(source io.Reader) (int, map[string]int, error) {
	var sue int
	var p1, p2, p3 string
	var v1, v2, v3 int
	_, err := fmt.Fscanf(
		source,
		"Sue %d: %s %d, %s %d, %s %d\n",
		&sue, &p1, &v1, &p2, &v2, &p3, &v3)
	if err != nil {
		return 0, nil, err
	}
	properties := map[string]int{
		strings.TrimRight(p1, ":"): v1,
		strings.TrimRight(p2, ":"): v2,
		strings.TrimRight(p3, ":"): v3,
	}
	return sue, properties, nil
}

func parse(input string) map[int]map[string]int {
	sues := make(map[int]map[string]int, 500)
	reader := strings.NewReader(input)
	for {
		if sue, properties, err := parseEntry(reader); err == nil {
			sues[sue] = properties
		} else if errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF) {
			// Getting ErrUnexpectedEOF here for some weird reason I can't discern.
			// Don't care: everything's being read.
			break
		} else {
			fmt.Printf("%v\n", err)
			panic(err)
		}
	}
	return sues
}
