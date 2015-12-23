package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"aoc/day6"
)

func main() {
	flag.Parse()

	var source io.Reader
	if flag.Arg(0) == "" {
		source = os.Stdin
	} else if file, err := os.Open(flag.Arg(0)); err != nil {
		log.Fatal(err)
	} else {
		defer file.Close()
		source = file
	}

	lights := day6.Lights{}

	day6.ParseReader(source, func(cmd string, flag bool, from, to day6.Coord) {
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
