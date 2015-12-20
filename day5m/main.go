package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"aoc/day5"
)

func main() {
	flag.Parse()

	path := flag.Arg(0)
	if path == "" {
		log.Fatal("Provide at least one input file path")
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	nNice := 0
	nTotal := 0
	for scanner.Scan() {
		nTotal++
		if day5.IsNice1(scanner.Text()) {
			nNice++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Nice entries: %v out of %v\n", nNice, nTotal)
}
