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

type trump struct {
	choice, beats rune
}

var beats []trump

func scorePlay(p1, p2 rune) int {
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
	fmt.Printf("Part 1 score: %v\n", score)
}

func getCorrectPlay(p1, hint rune) rune {
	switch hint {
	case 'X': // Lose
		for _, pairwise := range beats {
			if p1 == pairwise.choice {
				return pairwise.beats
			}
		}
	case 'Z': // Win
		for _, pairwise := range beats {
			if p1 == pairwise.beats {
				return pairwise.choice
			}
		}
	}
	// Default: draw
	return p1
}

func part2(guide []pair) {
	score := 0
	for _, entry := range guide {
		attempt := getCorrectPlay(entry.p1, entry.p2)
		score += scorePlay(entry.p1, attempt)
	}
	fmt.Printf("Part 2 score: %v\n", score)
}

func init() {
	beats = []trump{
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
}

func main() {
	input := readInput()
	part1(input)
	part2(input)
}
