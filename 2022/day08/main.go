package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Plot struct {
	Height int
	// Directions from which this tree is visible. Assume visibility by default.
	Left, Right, Top, Bottom bool
}

func (v Plot) IsVisible() bool {
	return v.Left || v.Right || v.Top || v.Bottom
}

func (v *Plot) Set(height int) {
	v.Height = height
	v.Top = true
	v.Bottom = true
	v.Left = true
	v.Right = true
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	grid := make([][]Plot, 0)
	scanner := bufio.NewScanner(f)
	for y := 0; scanner.Scan(); y++ {
		runes := []rune(scanner.Text())
		grid = append(grid, make([]Plot, len(runes)))
		for x := 0; x < len(runes); x++ {
			(&grid[y][x]).Set(int(runes[x] - '0'))
		}
	}

	// Part 1: scan the grid, knocking out any invisible trees

	height := len(grid)
	width := len(grid[0])

	for y := 0; y < height; y++ {
		// -1 is going to be greater than anything at the edge, so it's a good starting point
		leftHighest := -1
		rightHighest := -1
		for x := 0; x < width; x++ {
			if grid[y][x].Height <= leftHighest {
				(&grid[y][x]).Left = false
			} else {
				leftHighest = grid[y][x].Height
			}
			xOpposite := width - x - 1
			if grid[y][xOpposite].Height <= rightHighest {
				(&grid[y][xOpposite]).Right = false
			} else {
				rightHighest = grid[y][xOpposite].Height
			}
		}
	}

	for x := 0; x < width; x++ {
		// -1 is going to be greater than anything at the edge, so it's a good starting point
		topHighest := -1
		bottomHighest := -1
		for y := 0; y < height; y++ {
			if grid[y][x].Height <= topHighest {
				(&grid[y][x]).Top = false
			} else {
				topHighest = grid[y][x].Height
			}
			yOpposite := height - y - 1
			if grid[yOpposite][x].Height <= bottomHighest {
				(&grid[yOpposite][x]).Bottom = false
			} else {
				bottomHighest = grid[yOpposite][x].Height
			}
		}
	}

	nVisible := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x].IsVisible() {
				nVisible++
			}
		}
	}

	fmt.Printf("Part 1: %v visible\n", nVisible)
}
