package main

import (
	"fmt"
	"os"

	"github.com/kgaughan/aoc/2015/day1/lib"
)

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Printf(
			"%q is floor %v; basement entered at %v\n",
			arg,
			lib.CountFloors(arg),
			lib.FindBasementInstruction(arg))
	}
}
