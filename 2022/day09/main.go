package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
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

func (m move) String() string {
	return fmt.Sprintf("{%c %v}", m.direction, m.distance)
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

func (r *rope) String() string {
	maxX, maxY := 0, 0
	minX, minY := 0, 0
	for _, knot := range r.knots {
		maxX = max(knot.x, maxX)
		maxY = max(knot.y, maxY)
		minX = min(knot.x, minX)
		minY = min(knot.y, minY)
	}
	for visited := range r.TailTrail {
		maxX = max(visited.x, maxX)
		maxY = max(visited.y, maxY)
		minX = min(visited.x, minX)
		minY = min(visited.y, minY)
	}
	result := &strings.Builder{}
	for y := maxY; y >= minY; y-- {
	line:
		for x := minX; x <= maxX; x++ {
			for i, knot := range r.knots {
				if knot.x == x && knot.y == y {
					if i == 0 {
						result.WriteRune('H')
					} else {
						fmt.Fprintf(result, "%d", i)
					}
					continue line
				}
			}
			if _, exists := r.TailTrail[coordinate{x: x, y: y}]; exists {
				result.WriteRune('#')
			} else {
				result.WriteRune('.')
			}
		}
		result.WriteRune('\n')
	}
	return result.String()
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
