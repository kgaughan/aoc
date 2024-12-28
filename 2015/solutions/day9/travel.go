package day9

import (
	"math"

	"github.com/kgaughan/aoc/2015/helpers/permute"
)

func travel(start int, nodes []string, edges Edges, cond func(int, int) bool) int {
	permute := permute.NewPermute(nodes)
	result := start
	for {
		permutation, done := permute.Get()
		if done {
			return result
		}

		if distance, ok := edges.Distance(permutation); ok {
			if cond(result, distance) {
				result = distance
			}
		}
	}
}

func TravelMin(nodes []string, edges Edges) int {
	return travel(math.MaxInt32, nodes, edges, func(current, next int) bool {
		return current > next
	})
}

func TravelMax(nodes []string, edges Edges) int {
	return travel(0, nodes, edges, func(current, next int) bool {
		return next > current
	})
}
