package day19

import (
	"container/heap"
	"fmt"
	"sort"
	"strings"

	"github.com/kgaughan/aoc/2015/helpers"
)

func Part2(input string) {
	reverseMappings := make(map[string]string, 50)
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
			reverseMappings[value] = key
		} else {
			molecule = s
		}
	})

	result := reduce(molecule, reverseMappings)
	fmt.Printf("%v\n", result)
}

func reduce(molecule string, reverseMappings map[string]string) int {
	replacements := make([]string, 0, len(reverseMappings))
	for replacement := range reverseMappings {
		replacements = append(replacements, replacement)
	}
	sort.Slice(replacements, func(i, j int) bool {
		return len(replacements[i]) > len(replacements[j])
	})

	seen := make(map[string]bool, 200)

	pq := PriorityQueue{}
	pq.Push(&Item{molecule: molecule, steps: 0})
	heap.Init(&pq)
	for {
		item := heap.Pop(&pq).(*Item)
		for _, replacement := range replacements {
			key := reverseMappings[replacement]
			offsets := allIndexes(item.molecule, replacement)
			for _, offset := range offsets {
				prefix := item.molecule[:offset]
				suffix := item.molecule[offset+len(replacement):]
				newMolecule := prefix + key + suffix
				if _, ok := seen[newMolecule]; ok {
					continue
				}
				seen[newMolecule] = true
				if newMolecule == "e" {
					return item.steps + 1
				}
				heap.Push(&pq, &Item{
					molecule: newMolecule,
					steps:    item.steps + 1,
				})
			}
		}
	}
}
