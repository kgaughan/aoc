package day9

import (
	"fmt"
)

func Part2(input string) {
	edges, towns := parse(input)
	distance := TravelMax(towns, edges)
	fmt.Printf("Longest distance is %v\n", distance)
}
