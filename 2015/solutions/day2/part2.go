package day2

import "fmt"

func Part2(input string) {
	dimensions := parse(input)
	var total int
	for _, d := range dimensions {
		total += d.ribbon()
	}
	fmt.Printf("total ribbon required is %v\n", total)
}
