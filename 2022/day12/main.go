package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type coordinate struct {
	x, y int
}

func (c coordinate) canVisit(current coordinate, terrain [][]rune, visited map[coordinate]bool) bool {
	// Check bounds
	if c.x < 0 || c.y < 0 || c.x >= len(terrain[0]) || c.y >= len(terrain) {
		return false
	}
	// Have we been here?
	if _, exists := visited[c]; exists {
		return false
	}
	// We can only go up one higher, but we can drop down an arbitrary amount.
	return terrain[c.y][c.x] <= terrain[current.y][current.x]+1
}

func findPathLength(terrain [][]rune, start, end coordinate) int {
	visited := make(map[coordinate]bool)
	visited[start] = true
	next := make([]coordinate, 0, 10)
	next = append(next, start)
	distance := 0
	for len(next) > 0 {
		queued := make([]coordinate, 0, len(next))
		for _, here := range next {
			if here == end {
				return distance
			}

			up := coordinate{x: here.x, y: here.y - 1}
			if up.canVisit(here, terrain, visited) {
				queued = append(queued, up)
				visited[up] = true
			}

			down := coordinate{x: here.x, y: here.y + 1}
			if down.canVisit(here, terrain, visited) {
				queued = append(queued, down)
				visited[down] = true
			}

			left := coordinate{x: here.x - 1, y: here.y}
			if left.canVisit(here, terrain, visited) {
				queued = append(queued, left)
				visited[left] = true
			}

			right := coordinate{x: here.x + 1, y: here.y}
			if right.canVisit(here, terrain, visited) {
				queued = append(queued, right)
				visited[right] = true
			}
		}
		next = queued
		distance++
	}
	return -1
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	terrain := make([][]rune, 0)
	scanner := bufio.NewScanner(f)
	start := coordinate{x: 0, y: -1}
	end := coordinate{x: 0, y: -1}
	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		terrain = append(terrain, []rune(line))
		if start.y == -1 {
			start.x = strings.IndexRune(line, 'S')
			if start.x != -1 {
				start.y = y
				terrain[start.y][start.x] = 'a'
			}
		}
		if end.y == -1 {
			end.x = strings.IndexRune(line, 'E')
			if end.x != -1 {
				end.y = y
				terrain[end.y][end.x] = 'z'
			}
		}
	}

	fmt.Printf("Part 1: %v\n", findPathLength(terrain, start, end))

	shortestHike := -1
	for y := 0; y < len(terrain); y++ {
		for x, cell := range terrain[y] {
			if cell == 'a' {
				distance := findPathLength(terrain, coordinate{x: x, y: y}, end)
				if distance > 0 && (shortestHike == -1 || distance < shortestHike) {
					shortestHike = distance
				}
			}
		}
	}
	fmt.Printf("Part 2: %v\n", shortestHike)
}
