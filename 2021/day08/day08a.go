package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

type runeString []rune

func (s runeString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s runeString) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s runeString) Len() int {
	return len(s)
}

func sortString(in string) string {
	runes := []rune(in)
	sort.Sort(runeString(runes))
	return string(runes)
}

func parseLine(r io.Reader) ([]string, []string, error) {
	signals := make([]string, 10)
	reading := make([]string, 4)
	_, err := fmt.Fscanf(
		r, "%s %s %s %s %s %s %s %s %s %s | %s %s %s %s\n",
		&signals[0],
		&signals[1],
		&signals[2],
		&signals[3],
		&signals[4],
		&signals[5],
		&signals[6],
		&signals[7],
		&signals[8],
		&signals[9],
		&reading[0],
		&reading[1],
		&reading[2],
		&reading[3],
	)
	if err != nil {
		return nil, nil, err
	}
	// Sorting the runes makes matching 1, 4, 7, and 8 a bit easier for part 1,
	// but will likely be unnecessary for part 2, aside from aiding debugging.
	for i, s := range signals {
		signals[i] = sortString(s)
	}
	for i, s := range reading {
		reading[i] = sortString(s)
	}
	return signals, reading, nil
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	occurrences := 0
	for {
		signals, reading, err := parseLine(f)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(signals, reading)
		digits := make(map[string]int, 10)
		for _, s := range signals {
			if len(s) == 2 {
				digits[s] = 1
			} else if len(s) == 4 {
				digits[s] = 4
			} else if len(s) == 3 {
				digits[s] = 7
			} else if len(s) == 7 {
				digits[s] = 8
			}
		}
		for _, s := range reading {
			if _, ok := digits[s]; ok {
				occurrences++
			}
		}
	}

	fmt.Printf("Occurrences of 1, 4, 7, or 8: %d\n", occurrences)
}
