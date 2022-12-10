package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type move struct {
	direction rune
	distance  int
}

func (m move) ToOffset() (int, int) {
	dx := 0
	dy := 0
	switch m.direction {
	case 'L':
		dx = -m.distance
	case 'R':
		dx = m.distance
	case 'U':
		dy = m.distance
	case 'D':
		dy = -m.distance
	}
	return dx, dy
}

type coordinate struct {
	x, y int
}

func (c coordinate) GetOffset(other coordinate) (int, int) {
	return c.x - other.x, c.y - other.y
}

type rope struct {
	head, tail coordinate
	TailTrail  map[coordinate]bool
}

func (r *rope) Move(move move) {
	dx, dy := move.ToOffset()
	r.head.x += dx
	r.head.y += dy
	mdx, mdy := r.head.GetOffset(r.tail)
	for abs(mdx) > 1 || abs(mdy) > 1 {
		xMove, yMove := sgn(mdx), sgn(mdy)
		r.tail.x += xMove
		r.tail.y += yMove
		mdx -= xMove
		mdy -= yMove
		r.TailTrail[r.tail] = true
	}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func sgn(n int) int {
	switch {
	case n < 0:
		return -1
	case n > 0:
		return 1
	}
	return 0
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	rope := &rope{
		TailTrail: make(map[coordinate]bool),
	}
	rope.TailTrail[coordinate{x: 0, y: 0}] = true

	for {
		var move move
		_, err := fmt.Fscanf(f, "%c %d\n", &move.direction, &move.distance)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		rope.Move(move)
	}
	fmt.Printf("Part 1: %v unique visits\n", len(rope.TailTrail))
}
