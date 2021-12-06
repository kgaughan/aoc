package bingo

import (
	"strings"
	"testing"
)

// TestRead ensures grids are read correctly.
func TestRead(t *testing.T) {
	r := strings.NewReader(`22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19
`)
	g := Grid{}
	if err := g.Read(r); err != nil {
		t.Error(err)
	}
	// Check some key numbers in the grid
	expectations := []struct {
		x, y, value int
	}{
		{x: 0, y: 0, value: 22},
		{x: 4, y: 0, value: 0},
		{x: 0, y: 1, value: 8},
		{x: 4, y: 3, value: 5},
		{x: 4, y: 4, value: 19},
	}
	for _, e := range expectations {
		if g.numbers[e.y][e.x] != e.value {
			t.Errorf(
				"At (%d,%d), expected %d, but found %d",
				e.x, e.y,
				e.value,
				g.numbers[e.y][e.x],
			)
		}
	}
}
