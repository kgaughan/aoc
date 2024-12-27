package day9

import (
	"fmt"

	"github.com/kgaughan/aoc/2015/helpers/graph"
)

func Part1(input string) {
	edges, towns := parse(input)
	distance := graph.TravelMin(towns, edges, false)
	fmt.Printf("Shortest distance is %v\n", distance)
}
