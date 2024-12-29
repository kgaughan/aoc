package day18

import "fmt"

func Part1(input string) {
	grid := readInput(input)
	for i := 0; i < 100; i++ {
		runAutomaton(grid, day18Offsets, day18Rule)
	}
	fmt.Printf("%v\n", countLights(grid))
}
