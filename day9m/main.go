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

	permute := day9.NewPermute(towns[1:])
	shortest := math.MaxInt32
	for {
		permutation, done := permute.Get()
		if done {
			break
		}
		distance := GetDistance(towns[0], permutation, pairs)
		if distance < shortest {
			shortest = distance
		}
	}
	fmt.Printf("Shortest distance is %v\n", shortest)
}

func GetDistance(start string, rest []string, pairs map[pair]int) int {
	total := pairs[pair{start, rest[0]}]
	for i := 0; i < len(rest)-1; i++ {
		total += pairs[pair{rest[i], rest[i+1]}]
	}
	return total
}
