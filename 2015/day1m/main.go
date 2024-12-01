package main

import (
	"aoc/day1"
	"fmt"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Printf(
			"%q is floor %v; basement entered at %v\n",
			arg,
			day1.CountFloors(arg),
			day1.FindBasementInstruction(arg))
	}
}
