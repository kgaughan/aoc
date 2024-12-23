package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/kgaughan/aoc/2015/day5/lib"
)

func main() {
	method2 := flag.Bool("2", false, "Use the second niceness method.")
	flag.Parse()

	var isNice lib.NiceFunc
	if *method2 {
		isNice = lib.IsNice2
	} else {
		isNice = lib.IsNice1
	}

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
		if isNice(scanner.Text()) {
			nNice++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Nice entries: %v out of %v\n", nNice, nTotal)
}
