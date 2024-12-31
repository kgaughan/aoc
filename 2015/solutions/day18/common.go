package day18

import (
	"strings"

	"github.com/kgaughan/aoc/2015/helpers"
)

func readInput(input string) [][]bool {
	rows := make([][]bool, 0, 100)
	helpers.ScanLines(strings.NewReader(input), func(s string) error {
		row := make([]bool, len(s))
		for i := 0; i < len(s); i++ {
			if s[i] == '#' {
				row[i] = true
			}
		}
		rows = append(rows, row)
		return nil
	})
	return rows
}

type point struct{ x, y int }

var day18Offsets = []point{
	{-1, -1},
	{0, -1},
	{1, -1},
	{-1, 0},
	{1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
}

func day18Rule(state bool, neighbours int) bool {
	if state {
		return neighbours == 2 || neighbours == 3
	}
	return neighbours == 3
}

func runAutomaton(grid [][]bool, offsets []point, rule func(state bool, neighbours int) bool) {
	width := len(grid[0])
	prev := make([]bool, width)
	scratch := make([]bool, width)
	for y := 0; y < len(grid); y++ {
		for x := 0; x < width; x++ {
			// count neighbours
			n := 0
			for _, offset := range offsets {
				oy := offset.y + y
				if oy < len(grid) {
					ox := offset.x + x
					if ox >= 0 && ox < width {
						row := prev // just to have something
						if offset.y >= 0 {
							row = grid[y+offset.y]
						}
						if row[ox] {
							n++
						}
					}
				}
			}
			scratch[x] = rule(grid[y][x], n)
		}

		// duplicate this lines for the next iteration
		copy(prev, grid[y])
		// overwrite with the new state
		copy(grid[y], scratch)
	}
}

func countLights(grid [][]bool) int {
	n := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell {
				n++
			}
		}
	}
	return n
}
