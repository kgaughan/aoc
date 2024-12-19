package main

import (
	"fmt"
	"io"
	"log"

	"github.com/kgaughan/aoc/2015/day2/lib"
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
		area := lib.Area(l, w, h)
		ribbon := lib.Ribbon(l, w, h)
		fmt.Printf("Area for %vx%vx%v is %v; with a ribbon of length %v\n", l, w, h, area, ribbon)
		totalArea += area
		totalRibbon += ribbon
	}
	fmt.Printf("Total area required is %v; total ribbon required is %v\n", totalArea, totalRibbon)
}
