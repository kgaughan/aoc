package day3

import (
	"flag"
	"fmt"
)

var workers = flag.Int("day3workers", 1, "Number of workers to use for day 3")

func Part1(input string) {
	visited := deliver(input, *workers)
	fmt.Printf("Houses visited: %v\n", visited)
}
