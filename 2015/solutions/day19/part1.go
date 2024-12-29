package day19

import (
	"fmt"
	"strings"

	"github.com/kgaughan/aoc/2015/helpers"
)

func Part1(input string) {
	mappings := make(map[string][]string, 50)
	molecule := ""
	parsingMapping := true
	helpers.ScanLines(strings.NewReader(input), func(s string) {
		if s == "" {
			parsingMapping = false
		} else if parsingMapping {
			var key, value string
			if _, err := fmt.Sscanf(s, "%s => %s\n", &key, &value); err != nil {
				panic(err)
			}
			mappings[key] = append(mappings[key], value)
		} else {
			molecule = s
		}
	})

	replacements := make(map[string]bool, 100)
	for k, vs := range mappings {
		offsets := allIndexes(molecule, k)
		for _, offset := range offsets {
			prefix := molecule[:offset]
			suffix := molecule[offset+len(k):]
			for _, v := range vs {
				replacements[prefix+v+suffix] = true
			}
		}
	}

	fmt.Printf("%v\n", len(replacements))
}
