package main

import (
	"fmt"
	"io"
	"log"

	"aoc/day2"
)

func main() {
	totalArea := 0
	totalRibbon := 0
	var l, w, h int
	for {
		if _, err := fmt.Scanf("%dx%dx%d", &l, &w, &h); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		area := day2.Area(l, w, h)
		ribbon := day2.Ribbon(l, w, h)
		fmt.Printf("Area for %vx%vx%v is %v; with a ribbon of length %v\n", l, w, h, area, ribbon)
		totalArea += area
		totalRibbon += ribbon
	}
	fmt.Printf("Total area required is %v; total ribbon required is %v\n", totalArea, totalRibbon)
}
