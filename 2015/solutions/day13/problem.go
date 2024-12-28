package day13

import "github.com/kgaughan/aoc/2015/helpers/permute"

type pair struct{ person, other string }

func tryAll(entries map[pair]int, addSelf bool) int {
	people := make(map[string]bool, 10)
	for pair := range entries {
		people[pair.person] = true
	}
	keys := make([]string, 0, len(people)+1)
	for key := range people {
		keys = append(keys, key)
	}

	if addSelf {
		for _, key := range keys {
			entries[pair{person: key, other: "Self"}] = 0
			entries[pair{person: "Self", other: key}] = 0
		}
		keys = append(keys, "Self")
	}

	permute := permute.NewPermute(keys)
	best := 0
	for {
		permutation, done := permute.Get()
		if done {
			return best
		}
		happiness := 0
		for i, person := range permutation {
			j := i - 1
			if j < 0 {
				j += len(permutation)
			}
			other := permutation[j]
			happiness += entries[pair{person: person, other: other}] + entries[pair{person: other, other: person}]
		}
		if happiness > best {
			best = happiness
		}
	}
}
