package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/kgaughan/aoc/2015/day7/lib"
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

	lib.Parse(source)
}
