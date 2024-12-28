package day16

import "fmt"

func Part2(input string) {
	sues := parse(input)
	matched := findSue(sues, func(key string, value, expected int) bool {
		switch key {
		case "trees", "cats":
			return value > expected
		case "pomeranians", "goldfish":
			return value < expected
		}
		return value == expected
	})
	fmt.Printf("%v\n", matched)
}
