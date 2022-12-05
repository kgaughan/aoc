package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type span struct {
	start, end int
}

func (s span) Contains(other span) bool {
	return s.start >= other.start && s.end <= other.end
}

type pair struct {
	s1, s2 span
}

func readInput() []pair {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	result := make([]pair, 0, 10)

	for {
		var entry pair
		_, err := fmt.Fscanf(f, "%d-%d,%d-%d\n", &entry.s1.start, &entry.s1.end, &entry.s2.start, &entry.s2.end)
		if err == io.EOF {
			break
		}
		result = append(result, entry)
	}
	return result
}

func part1(input []pair) {
	n := 0
	for _, entry := range input {
		if entry.s1.Contains(entry.s2) || entry.s2.Contains(entry.s1) {
			n++
		}
	}
	fmt.Printf("Part 1: %v\n", n)
}

func part2(input []pair) {
}

func main() {
	input := readInput()
	part1(input)
	part2(input)
}
