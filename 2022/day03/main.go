package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type rucksack struct {
	compartment1, compartment2 string
}

func (sack rucksack) all() string {
	return sack.compartment1 + sack.compartment2
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func readInput() []rucksack {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	result := make([]rucksack, 0, 10)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		length := len(line) / 2
		result = append(result, rucksack{
			compartment1: line[:length],
			compartment2: line[length:],
		})
	}

	return result
}

func findCommon(items ...string) []rune {
	common := make([]rune, 0, 4)
	indices := make([]int, len(items))
	runeSets := make([][]rune, len(items))
	for i, item := range items {
		runes := []rune(item)
		sort.Sort(sortRunes(runes))
		indices[i] = 0
		runeSets[i] = runes
	}
loop:
	for {
		// If we've reached the end of any set, we're done
		for i, index := range indices {
			if index >= len(runeSets[i]) {
				break loop
			}
		}

		// Check if all the current elements are the same
		same := true
		candidate := runeSets[0][indices[0]]
		lowest := candidate
		for i, rs := range runeSets {
			if rs[indices[i]] != candidate {
				// This scan didn't find a good candidate
				same = false
			}
			// Find the lowest rune: we increment the indices of anything with
			// this later
			if rs[indices[i]] < lowest {
				lowest = rs[indices[i]]
			}
		}

		// We have a good candidate, so append it
		if same && (len(common) == 0 || common[len(common)-1] != candidate) {
			common = append(common, candidate)
		}

		// Move forward where necessary
		for i, rs := range runeSets {
			for indices[i] < len(runeSets[i]) && rs[indices[i]] == lowest {
				indices[i]++
			}
		}
	}

	return common
}

func getPriority(ch rune) int {
	if ch >= 'a' && ch <= 'z' {
		return int(ch-'a') + 1
	}
	if ch >= 'A' && ch <= 'Z' {
		return int(ch-'A') + 27
	}
	return 0
}

func part1(input []rucksack) {
	score := 0
	for _, sack := range input {
		for _, ch := range findCommon(sack.compartment1, sack.compartment2) {
			score += getPriority(ch)
		}
	}
	fmt.Printf("Part 1, sum of priorities: %v\n", score)
}

func part2(input []rucksack) {
	score := 0
	for i := 0; i < len(input); i += 3 {
		for _, ch := range findCommon(input[i].all(), input[i+1].all(), input[i+2].all()) {
			score += getPriority(ch)
		}
	}
	fmt.Printf("Part 2, sum of priorities: %v\n", score)

}

func main() {
	input := readInput()
	part1(input)
	part2(input)
}
