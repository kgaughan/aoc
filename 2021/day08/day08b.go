package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

const signals = "abcdefg"

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
	return signals, reading, nil
}

func fact(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func permute(chars string) []string {
	result := make([]string, 0, fact(len(chars)))
	generate(&result, []rune(chars), 0, len(chars)-1)
	return result
}

func generate(result *[]string, acc []rune, left, right int) {
	if left == right {
		*result = append(*result, string(acc))
	} else {
		for i := left; i <= right; i++ {
			acc[left], acc[i] = acc[i], acc[left]
			generate(result, acc, left+1, right)
			acc[left], acc[i] = acc[i], acc[left]
		}
	}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	permutations := permute(signals)
	correctPositions := []string{
		"abcefg",  // 0
		"cf",      // 1
		"acdeg",   // 2
		"acdfg",   // 3
		"bcdf",    // 4
		"abdfg",   // 5
		"abdefg",  // 6
		"acf",     // 7
		"abcdefg", // 8
		"abcdfg",  // 9
	}

	for {
		signals, reading, err := parseLine(f)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(signals, reading)
	}
}
