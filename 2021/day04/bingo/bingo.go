package bingo

import (
	"fmt"
	"io"
)

// Grid is a representation of a bingo grid.
type Grid struct {
	numbers [5][5]int
	marked  [5][5]bool
}

// Read reads five lines from the given reader into the grid.
func (g *Grid) Read(r io.Reader) error {
	for y := 0; y < len(g.numbers); y++ {
		_, err := fmt.Fscanf(
			r, "%d %d %d %d %d\n",
			&g.numbers[y][0],
			&g.numbers[y][1],
			&g.numbers[y][2],
			&g.numbers[y][3],
			&g.numbers[y][4],
		)
		for x := 0; x < len(g.marked[y]); x++ {
			g.marked[y][x] = false
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Grid) Mark(n int) {
	for y := 0; y < len(g.numbers); y++ {
		for x := 0; x < len(g.numbers[y]); x++ {
			if g.numbers[y][x] == n {
				g.marked[y][x] = true
				return
			}
		}
	}
}

func (g *Grid) IsWinning() bool {
	// Scan rows
	for y := 0; y < len(g.numbers); y++ {
		win := true
		for x := 0; x < len(g.numbers[y]); x++ {
			if !g.marked[y][x] {
				win = false
				break
			}
		}
		if win {
			return true
		}
	}

	// Scan columns
	for x := 0; x < len(g.numbers[0]); x++ {
		win := true
		for y := 0; y < len(g.numbers); y++ {
			if !g.marked[y][x] {
				win = false
				break
			}
		}
		if win {
			return true
		}
	}

	// No wins
	return false
}

func (g *Grid) GetScore() int {
	sum := 0
	for y := 0; y < len(g.numbers); y++ {
		for x := 0; x < len(g.numbers[y]); x++ {
			if !g.marked[y][x] {
				sum += g.numbers[y][x]
			}
		}
	}
	return sum
}
