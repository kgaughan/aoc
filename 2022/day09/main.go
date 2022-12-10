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
	knots     []coordinate
	TailTrail map[coordinate]bool
}

func NewRope(length int) *rope {
	rope := &rope{
		knots:     make([]coordinate, length),
		TailTrail: make(map[coordinate]bool),
	}
	rope.TailTrail[coordinate{x: 0, y: 0}] = true
	return rope
}

func (r *rope) Move(move move) {
	hdx, hdy := move.ToOffset()
	for abs(hdx) > 0 || abs(hdy) > 0 {
		xMove, yMove := sgn(hdx), sgn(hdy)
		r.knots[0].x += xMove
		r.knots[0].y += yMove
		hdx -= xMove
		hdy -= yMove
		for i := 1; i < len(r.knots); i++ {
			mdx, mdy := r.knots[i-1].GetOffset(r.knots[i])
			for abs(mdx) > 1 || abs(mdy) > 1 {
				xMove, yMove = sgn(mdx), sgn(mdy)
				r.knots[i].x += xMove
				r.knots[i].y += yMove
				mdx -= xMove
				mdy -= yMove
				if i == len(r.knots)-1 {
					r.TailTrail[r.knots[len(r.knots)-1]] = true
				}
			}
		}
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

	rope1 := NewRope(2)
	rope2 := NewRope(10)

	for {
		var move move
		_, err := fmt.Fscanf(f, "%c %d\n", &move.direction, &move.distance)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		rope1.Move(move)
		rope2.Move(move)
	}
	fmt.Printf("Part 1: %v unique visits\n", len(rope1.TailTrail))
	fmt.Printf("Part 2: %v unique visits\n", len(rope2.TailTrail))
}
