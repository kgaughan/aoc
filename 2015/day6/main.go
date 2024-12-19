package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/kgaughan/aoc/2015/day6/lib"
)

func main() {
	method2 := flag.Bool("2", false, "Use method 2.")
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

	lights := lib.Lights{}

	var method lib.CmdFunc
	if !*method2 {
		method = func(cmd string, flag bool, from, to lib.Coord) {
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
		}
	} else {
		method = func(cmd string, flag bool, from, to lib.Coord) {
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
		}
	}

	lib.ParseReader(source, method)

	fmt.Printf("Total number of lights on is: %v\n", lights.Count())
}
