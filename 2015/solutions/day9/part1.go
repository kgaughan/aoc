package day9

import (
	"fmt"

	"github.com/kgaughan/aoc/2015/helpers/graph"
)

func Part2(input string) {
	edges, towns := parse(input)
	distance := graph.TravelMax(towns, edges, false)
	fmt.Printf("Longest distance is %v\n", distance)
}
