package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

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
		// Vertical and horizontal line can be treated as boxes.
		if x1 == x2 || y1 == y2 {
			if x1 > x2 {
				x1, x2 = x2, x1
			}
			if y1 > y2 {
				y1, y2 = y2, y1
			}
			for y := y1; y <= y2; y++ {
				if points[y] == nil {
					points[y] = make(map[int]int, 10)
				}
				for x := x1; x <= x2; x++ {
					points[y][x]++
					if points[y][x] == 2 {
						overlaps++
					}
				}
			}
		}
	}
	fmt.Printf("Overlapping points: %d\n", overlaps)
}
