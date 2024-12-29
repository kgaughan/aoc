package day18

import "fmt"

func Part2(input string) {
	grid := readInput(input)
	setCorners(grid)
	for i := 0; i < 100; i++ {
		runAutomaton(grid, day18Offsets, day18Rule)
		setCorners(grid)
	}
	fmt.Printf("%v\n", countLights(grid))
}

func setCorners(grid [][]bool) {
	width := len(grid[0])
	height := len(grid)
	grid[0][0] = true
	grid[height-1][0] = true
	grid[0][width-1] = true
	grid[height-1][width-1] = true
}
