package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type criterion func(int, int) bool

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	totals := make([]int, 0, 256)
	lines := make([]int64, 0, 2000)
	for scanner.Scan() {
		line := scanner.Text()
		for idx, ch := range line {
			if len(totals) <= idx {
				totals = append(totals, 0)
			}
			if ch == '1' {
				totals[idx] += 1
			}
		}
		parsed, err := strconv.ParseInt(line, 2, 64)
		if err != nil {
			log.Fatal(err)
		}
		lines = append(lines, parsed)
	}

	// Calculate fuel consumption
	gamma := 0
	epsilon := 0
	for _, n := range totals {
		gamma <<= 1
		epsilon <<= 1
		if n > len(lines)/2 {
			gamma |= 1
		} else {
			epsilon |= 1
		}
	}
	fmt.Printf("Power consumption: %v\n", gamma*epsilon)

	highBit := len(totals) - 1
	generatorRating := filter(lines, highBit, func(ones, scanLimit int) bool {
		return ones*2 >= scanLimit
	})
	scrubberRating := filter(lines, highBit, func(ones, scanLimit int) bool {
		return ones*2 < scanLimit
	})
	fmt.Printf("Life support rating: %v\n", generatorRating*scrubberRating)
}

func filter(lines []int64, highBit int, fn criterion) int64 {
	data := append([]int64{}, lines...)
	scanLimit := len(data)
	for i := highBit; i >= 0; i-- {
		ones := 0
		for j := 0; j < scanLimit; j++ {
			if data[j]&(1<<i) != 0 {
				ones++
			}
		}
		filterSet := fn(ones, scanLimit)
		newLimit := 0
		for j := 0; j < scanLimit; j++ {
			isSet := data[j]&(1<<i) > 0
			if isSet == filterSet {
				data[newLimit] = data[j]
				newLimit++
			}
		}
		scanLimit = newLimit
		if scanLimit == 1 {
			break
		}
	}
	return data[0]
}
