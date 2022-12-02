package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type pair struct {
	p1, p2 rune
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
		_, err := fmt.Fscanf(f, "%c %c\n", &entry.p1, &entry.p2)
		if err == io.EOF {
			break
		}
		result = append(result, entry)
	}

	return result
}

const (
	rock     = 'A'
	paper    = 'B'
	scissors = 'C'
)

const p2Offset = 23

func scorePlay(p1, p2 rune) int {
	beats := []struct {
		choice, beats rune
	}{
		{
			choice: rock,
			beats:  scissors,
		},
		{
			choice: scissors,
			beats:  paper,
		},
		{
			choice: paper,
			beats:  rock,
		},
	}

	score := 0
	switch p2 {
	case rock:
		score = 1
	case paper:
		score = 2
	case scissors:
		score = 3
	}

	var extra int
	for _, pairwise := range beats {
		if p2 == pairwise.choice {
			if p1 == pairwise.beats {
				extra = 6
			} else if p1 != p2 {
				extra = 0
			} else {
				extra = 3
			}
			break
		}
	}

	return score + extra
}

func part1(guide []pair) {
	score := 0
	for _, entry := range guide {
		score += scorePlay(entry.p1, entry.p2-p2Offset)
	}
	fmt.Printf("Cumulative score: %v\n", score)
}

func part2(guide []pair) {
}

func main() {
	input := readInput()
	part1(input)
	part2(input)
}
