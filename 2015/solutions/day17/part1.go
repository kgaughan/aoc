package day17

import "fmt"

func Part1(input string) {
	containers := readInput(input)
	result := combinations(containers, len(containers)-1, eggnog)
	fmt.Printf("%v\n", result)
}

func combinations(containers []int, iFinal, target int) int {
	if target == 0 {
		return 1
	}
	if iFinal < 0 {
		return 0
	}

	result := combinations(containers, iFinal-1, target)
	if containers[iFinal] <= target {
		result += combinations(containers, iFinal-1, target-containers[iFinal])
	}
	return result
}
