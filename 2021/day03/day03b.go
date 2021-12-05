package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

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

	// Calculate O2 generation. We start with a full list of values, and for
	// each bit from most to least significant, we determine the most common
	// bit in that column, then go over all the data a second time, discarding
	// any values that don't have that bit set.
	o2Data := append([]int64{}, lines...)
	o2ScanLimit := len(o2Data)
	for i := highBit; i >= 0; i-- {
		ones := 0
		for j := 0; j < o2ScanLimit; j++ {
			if o2Data[j]&(1<<i) != 0 {
				ones++
			}
		}
		filterSet := ones*2 >= o2ScanLimit
		newLimit := 0
		for j := 0; j < o2ScanLimit; j++ {
			isSet := o2Data[j]&(1<<i) > 0
			if isSet == filterSet {
				o2Data[newLimit] = o2Data[j]
				newLimit++
			}
		}
		o2ScanLimit = newLimit
		if o2ScanLimit == 1 {
			break
		}
	}

	// Calculate CO2 generation. We start with a full list of values, and for
	// each bit from most to least significant, we determine the least common
	// bit in that column, then go over all the data a second time, discarding
	// any values that don't have that bit set.
	co2Data := append([]int64{}, lines...)
	co2ScanLimit := len(co2Data)
	for i := highBit; i >= 0; i-- {
		ones := 0
		for j := 0; j < co2ScanLimit; j++ {
			if co2Data[j]&(1<<i) != 0 {
				ones++
			}
		}
		filterSet := ones*2 < co2ScanLimit
		newLimit := 0
		for j := 0; j < co2ScanLimit; j++ {
			isSet := co2Data[j]&(1<<i) > 0
			if isSet == filterSet {
				co2Data[newLimit] = co2Data[j]
				newLimit++
			}
		}
		co2ScanLimit = newLimit
		if co2ScanLimit == 1 {
			break
		}
	}

	fmt.Printf("Life support rating: %v\n", co2Data[0]*o2Data[0])
}
