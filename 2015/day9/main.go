package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/kgaughan/aoc/2015/common/graph"
	"github.com/kgaughan/aoc/2015/day9/lib"
)

func main() {
	doLongest := flag.Bool("longest", false, "Calculate longest rather than shortest.")
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

	edges := graph.Edges{}
	towns := make([]string, 0, 50)
	lib.ScanLines(source, func(from, to string, distance int) {
		if !lib.Contains(towns, from) {
			towns = append(towns, from)
		}
		if !lib.Contains(towns, to) {
			towns = append(towns, to)
		}
		edges.Add(from, to, distance)
	})

	if *doLongest {
		distance := graph.TravelMax(towns, edges, false)
		fmt.Printf("Longest distance is %v\n", distance)
	} else {
		distance := graph.TravelMin(towns, edges, false)
		fmt.Printf("Shortest distance is %v\n", distance)
	}
}
