package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func sgn(n int) int {
	if n > 0 {
		return 1
	}
	if n < 0 {
		return -1
	}
	return 0
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

	overlaps := 0
	points := map[int]map[int]int{}
	for {
		var x1, y1, x2, y2 int
		if _, err := fmt.Fscanf(f, "%d,%d -> %d,%d\n", &x1, &y1, &x2, &y2); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		dx := sgn(x2 - x1)
		dy := sgn(y2 - y1)
		steps := max((x2-x1)*dx, (y2-y1)*dy)
		x := x1
		y := y1
		for i := 0; i <= steps; i++ {
			if points[y] == nil {
				points[y] = make(map[int]int, 10)
			}
			points[y][x]++
			if points[y][x] == 2 {
				overlaps++
			}
			x += dx
			y += dy
		}
	}
	fmt.Printf("Overlapping points: %d\n", overlaps)
}
