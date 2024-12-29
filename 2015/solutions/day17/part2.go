package day17

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Part2(input string) {
	containers := readInput(input)
	combosFound := make(map[string]int, 50)
	stack := make([]int, 0, len(containers))
	minCombinations(containers, len(containers)-1, eggnog, stack, combosFound)
	shortest := math.MaxInt32
	found := 0
	for _, n := range combosFound {
		if n == shortest {
			found++
		} else if n < shortest {
			shortest = n
			found = 1
		}
	}
	fmt.Printf("%v\n", found)
}

func join(l []int) string {
	asStrings := make([]string, len(l))
	for i, n := range l {
		asStrings[i] = strconv.Itoa(n)
	}
	return strings.Join(asStrings, ",")
}

func minCombinations(containers []int, iFinal, target int, stack []int, combosFound map[string]int) {
	if target == 0 {
		combosFound[join(stack)] = len(stack)
	}
	if iFinal >= 0 {
		minCombinations(containers, iFinal-1, target, stack, combosFound)
		// I'm abusing how slices work here somewhat!
		stack = append(stack, iFinal)
		minCombinations(containers, iFinal-1, target-containers[iFinal], stack, combosFound)
	}
}
