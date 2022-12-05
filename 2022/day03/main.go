package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type rucksack struct {
	compartment1, compartment2 string
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

func findCommon(sack rucksack) []rune {
	common := make([]rune, 0, 4)
	// The data is short enough that it makes little sense to sort the
	// contents, so On^2 comparisons is justifiable, I think.
	for _, ch1 := range sack.compartment1 {
		for _, ch2 := range sack.compartment2 {
			if ch1 == ch2 {
				add := true
				// Skip any duplicates
				for _, ch3 := range common {
					if ch1 == ch3 {
						add = false
						break
					}
				}
				if add {
					common = append(common, ch1)
				}
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
		for _, ch := range findCommon(sack) {
			score += getPriority(ch)
		}
	}
	fmt.Printf("Part 1, sum of priorities: %v\n", score)
}

func part2(input []rucksack) {
}

func main() {
	input := readInput()
	part1(input)
	part2(input)
}
