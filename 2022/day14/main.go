package main

import (
	"bufio"
	"flag"
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

var addFloor = flag.Bool("floor", false, "Add the floor for part 2")

func main() {
	flag.Parse()

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	shapes := make([][]coordinate, 0, 10)
	maxY := math.MinInt
	for scanner.Scan() {
		shape := strings.Split(scanner.Text(), " -> ")
		parsedShape := make([]coordinate, 0)
		for _, unparsed := range shape {
			parts := strings.Split(unparsed, ",")
			x, err := strconv.Atoi(parts[0])
			if err != nil {
				log.Fatal(err)
			}
			y, err := strconv.Atoi(parts[1])
			if err != nil {
				log.Fatal(err)
			}
			maxY = max(maxY, y)
			parsedShape = append(parsedShape, coordinate{x: x, y: y})
		}
		shapes = append(shapes, parsedShape)
	}

	if *addFloor {
		shapes = append(shapes, []coordinate{
			{x: 0, y: maxY + 2},
			{x: 1000, y: maxY + 2},
		})
	}

	// Adjust to give some buffer space to simplify bounds checks. Also needed
	// for adding the floor
	maxY += 2

	grid := make([][]int8, maxY+1)
	for y := 0; y < len(grid); y++ {
		grid[y] = make([]int8, 1001)
	}
	for _, shape := range shapes {
		for i := 1; i < len(shape); i++ {
			// Assumption: shapes consist of just horizontal and vertical lines
			if shape[i-1].x == shape[i].x {
				// Vertical line
				start := min(shape[i-1].y, shape[i].y)
				end := max(shape[i-1].y, shape[i].y)
				adjustedX := shape[i].x
				for y := start; y <= end; y++ {
					grid[y][adjustedX] = ROCK
				}
			} else {
				// Horizontal line
				start := min(shape[i-1].x, shape[i].x)
				end := max(shape[i-1].x, shape[i].x)
				for x := start; x <= end; x++ {
					grid[shape[i].y][x] = ROCK
				}
			}
		}
	}

	units := 0
simulation:
	for {
		px := 500
		py := 0

		// Check if the entrance is blocked
		if grid[py][px] != EMPTY {
			fmt.Println("Source blocked!")
			break
		}

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
				fmt.Println("Sand overflow!")
				break simulation
			}
		}
		units++
	}
	fmt.Printf("%v units\n", units)
}
