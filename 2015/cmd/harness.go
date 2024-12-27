package main

import (
	"flag"
	"io/ioutil"
	"log"
	"strings"

	"github.com/kgaughan/aoc/2015/solutions"
)

func main() {
	var day int
	var input string

	flag.IntVar(&day, "day", 1, "day to execute")
	flag.StringVar(&input, "input", "", "input for section; prefix '@' to load a file")
	flag.Parse()

	var actualInput string
	if !strings.HasPrefix(input, "@") {
		actualInput = input
	} else {
		path := (input)[1 : len(input)-1]
		if contents, err := ioutil.ReadFile(path); err != nil {
			log.Fatalf("could not load input from %q: %v", path, err)
		} else {
			actualInput = string(contents)
		}
	}

	solutions.Execute(day, actualInput)
}
