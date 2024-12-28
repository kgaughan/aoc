package day9

import (
	"fmt"
)

func Part1(input string) {
	edges, towns := parse(input)
	distance := TravelMin(towns, edges)
	fmt.Printf("Shortest distance is %v\n", distance)
}
