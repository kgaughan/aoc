package main

import (
	"aoc/day7"
	"flag"
	"io"
	"log"
	"os"
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

	day7.Parse(source)
}
