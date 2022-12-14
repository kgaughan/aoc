package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	EMPTY = iota
	SAND
	ROCK
)

type coordinate struct {
	x, y int
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	shapes := make([][]coordinate, 0, 10)
	// Useful for finding the bounds of the rocks to figure out where the
	// abyss is. Also good for figuring out how much memory we need.
	maxX, minX, maxY := math.MinInt, math.MaxInt, math.MinInt
	for scanner.Scan() {
		shape := strings.Split(scanner.Text(), " -> ")
		parsedShape := make([]coordinate, 0)
		for _, unparsed := range shape {
			parts := strings.Split(unparsed, ",")
			x, err := strconv.Atoi(parts[0])
			if err != nil {
				log.Fatal(err)
			}
			maxX = max(maxX, x)
			minX = min(minX, x)
			y, err := strconv.Atoi(parts[1])
			if err != nil {
				log.Fatal(err)
			}
			maxY = max(maxY, y)
			parsedShape = append(parsedShape, coordinate{x: x, y: y})
		}
		shapes = append(shapes, parsedShape)
	}

	// Adjust to give some buffer space to simplify bounds checks
	minX--
	maxX++
	maxY++

	grid := make([][]int8, maxY+1)
	for y := 0; y < len(grid); y++ {
		grid[y] = make([]int8, maxX-minX+1)
	}
	for _, shape := range shapes {
		for i := 1; i < len(shape); i++ {
			// Assumption: shapes consist of just horizontal and vertical lines
			if shape[i-1].x == shape[i].x {
				// Vertical line
				start := min(shape[i-1].y, shape[i].y)
				end := max(shape[i-1].y, shape[i].y)
				adjustedX := shape[i].x - minX
				for y := start; y <= end; y++ {
					grid[y][adjustedX] = ROCK
				}
			} else {
				// Horizontal line
				start := min(shape[i-1].x, shape[i].x) - minX
				end := max(shape[i-1].x, shape[i].x) - minX
				for x := start; x <= end; x++ {
					grid[shape[i].y][x] = ROCK
				}
			}
		}
	}

	units := 0
simulation:
	for {
		px := 500 - minX
		py := 0

		for {
			if grid[py+1][px] == EMPTY {
				// A unit of sand always falls down one step if possible.
				py++
			} else if grid[py+1][px-1] == EMPTY {
				// Left looks good
				py++
				px--
			} else if grid[py+1][px+1] == EMPTY {
				// Right looks good
				py++
				px++
			} else {
				// At rest: try next particle
				grid[py][px] = SAND
				break
			}
			// Has this grain slid off the edge?
			if py >= maxY {
				break simulation
			}
		}
		units++
	}
	fmt.Printf("Part 1: %v units\n", units)
}
