package graph

import (
	"aoc/common/permute"
	"math"
)

func travel(start int, nodes []string, edges Edges, closed bool, cond func(int, int) bool) int {
	permute := permute.NewPermute(nodes[1:])
	result := start
	for {
		permutation, done := permute.Get()
		if done {
			return result
		}

		distance := 0
		if closed {
			distance, _ = edges.Get(nodes[0], nodes[len(nodes)-1])
		}
		distance += edges.Distance(nodes[0], permutation)
		if cond(result, distance) {
			result = distance
		}
	}
}

func TravelMin(nodes []string, edges Edges, closed bool) int {
	return travel(math.MaxInt32, nodes, edges, closed, func(current, next int) bool {
		return current > next
	})
}

func TravelMax(nodes []string, edges Edges, closed bool) int {
	return travel(0, nodes, edges, closed, func(current, next int) bool {
		return next > current
	})
}
