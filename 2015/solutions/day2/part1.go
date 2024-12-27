package day2

import "fmt"

func Part1(input string) {
	dimensions := parse(input)
	var total int
	for _, d := range dimensions {
		total += d.area()
	}
	fmt.Printf("total area required is %v\n", total)
}
