package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/kgaughan/aoc/2015/day3/lib"
)

func main() {
	n := flag.Int("n", 1, "Number of workers to use")
	flag.Parse()

	filePath := flag.Arg(0)
	if filePath == "" {
		fmt.Println("You need to provide a route file.")
		return
	}

	var route string
	if bytes, err := ioutil.ReadFile(filePath); err != nil {
		log.Fatal(err)
	} else {
		route = string(bytes)
	}

	visited := lib.Deliver(route, *n)
	fmt.Printf("Houses visited: %v\n", visited)
}
