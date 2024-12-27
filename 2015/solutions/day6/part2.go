package day6

import (
	"fmt"
)

func Part2(input string) {
	lights := Lights{}

	ParseString(input, func(cmd string, flag bool, from, to Coord) {
		switch cmd {
		case "turn":
			if flag {
				lights.Increment(from, to, 1)
			} else {
				lights.Increment(from, to, -1)
			}
		case "toggle":
			lights.Increment(from, to, 2)
		}
	})

	fmt.Printf("Total number of lights on is: %v\n", lights.Count())
}
