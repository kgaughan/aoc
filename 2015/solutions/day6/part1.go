package day6

import (
	"fmt"
)

func Part1(input string) {
	lights := Lights{}

	ParseString(input, func(cmd string, flag bool, from, to Coord) {
		switch cmd {
		case "turn":
			if flag {
				lights.TurnOn(from, to)
			} else {
				lights.TurnOff(from, to)
			}
		case "toggle":
			lights.Toggle(from, to)
		}
	})

	fmt.Printf("Total number of lights on is: %v\n", lights.Count())
}
