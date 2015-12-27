package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"

	"aoc/day9"
)

type pair struct {
	from, to string
}

func main() {
	doLongest := flag.Bool("longest", false, "Calculate longest rather than shortest.")
	flag.Parse()

	var source io.Reader
	if flag.Arg(0) == "" {
		source = os.Stdin
	} else if file, err := os.Open(flag.Arg(0)); err != nil {
		log.Fatal(err)
	} else {
		defer file.Close()
		source = file
	}

	pairs := make(map[pair]int)
	towns := make([]string, 0, 50)
	day9.ScanLines(source, func(from, to string, distance int) {
		if !day9.Contains(towns, from) {
			towns = append(towns, from)
		}
		if !day9.Contains(towns, to) {
			towns = append(towns, to)
		}
		pairs[pair{from, to}] = distance
		pairs[pair{to, from}] = distance
	})

	if *doLongest {
		distance := Travel(0, towns, pairs, func(current, next int) bool {
			return next > current
		})
		fmt.Printf("Longest distance is %v\n", distance)
	} else {
		distance := Travel(math.MaxInt32, towns, pairs, func(current, next int) bool {
			return current > next
		})
		fmt.Printf("Shortest distance is %v\n", distance)
	}
}

func Travel(start int, towns []string, pairs map[pair]int, cond func(int, int) bool) int {
	permute := day9.NewPermute(towns[1:])
	result := start
	for {
		permutation, done := permute.Get()
		if done {
			return result
		}
		distance := GetDistance(towns[0], permutation, pairs)
		if cond(result, distance) {
			result = distance
		}
	}
}

func GetDistance(start string, rest []string, pairs map[pair]int) int {
	total := pairs[pair{start, rest[0]}]
	for i := 0; i < len(rest)-1; i++ {
		total += pairs[pair{rest[i], rest[i+1]}]
	}
	return total
}
