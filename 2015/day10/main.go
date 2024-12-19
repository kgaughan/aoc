package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/kgaughan/aoc/2015/day10/lib"
)

func main() {
	n := flag.Int("n", 40, "Number of iterations to perform")
	flag.Parse()

	initial := flag.Arg(0)
	if initial == "" {
		log.Fatal("You must provide an initial string")
	}

	convoluted := initial
	for i := 0; i < *n; i++ {
		convoluted = lib.LookAndSay(convoluted)
	}

	fmt.Printf("The length after %v convolutions is %v\n", *n, len(convoluted))
}
