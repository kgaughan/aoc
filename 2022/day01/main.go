package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var maxElf int64 = 0
	var currentTotal int64 = 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if currentTotal > maxElf {
				maxElf = currentTotal
			}
			currentTotal = 0
		} else {
			val, err := strconv.ParseInt(line, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			currentTotal += val
		}
	}
	if currentTotal > maxElf {
		maxElf = currentTotal
	}
	fmt.Printf("Elf with max calories has %v calories\n", maxElf)
}

type Calories []int64

func (a Calories) Len() int {
	return len(a)
}

func (a Calories) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Calories) Less(i, j int) bool {
	return a[i] > a[j]
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	elfTotals := make(Calories, 0, 10)
	var currentTotal int64 = 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			elfTotals = append(elfTotals, currentTotal)
			currentTotal = 0
		} else {
			val, err := strconv.ParseInt(line, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			currentTotal += val
		}
	}
	elfTotals = append(elfTotals, currentTotal)
	sort.Sort(elfTotals)

	fmt.Printf("Total for top three elves is %v calories\n", elfTotals[0]+elfTotals[1]+elfTotals[2])
}

func main() {
	part1()
	part2()
}
