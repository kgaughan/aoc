package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/kgaughan/aoc/2015/solutions"
)

func main() {
	var day int
	var input string
	var inputPath string

	flag.IntVar(&day, "day", 1, "day to execute")
	flag.StringVar(&input, "input", "", "input for section if no file is given or found")
	flag.StringVar(&inputPath, "path", "", "use the given file as input; will attempt to infer a file if non specified and no input is given")
	flag.Parse()

	if input == "" {
		fallback := fmt.Sprintf("input/day%d.txt", day)
		if inputPath == "" && isFile(fallback) {
			inputPath = fallback
		}
		if inputPath == "" {
			log.Fatalf("no input provided")
		}
		contents, err := ioutil.ReadFile(inputPath)
		if err != nil {
			log.Fatalf("could not load input from %q: %v", inputPath, err)
		}
		input = string(contents)
	}
	solutions.Execute(day, input)
}

func isFile(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !fi.IsDir()
}
