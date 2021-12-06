package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/kgaughan/aoc/2021/day04/bingo"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	numbers, err := bingo.ReadNumbers(reader)
	if err != nil {
		log.Fatal(err)
	}

	grids, err := bingo.ReadGrids(reader)
	if err != nil {
		log.Fatal(err)
	}

	for _, n := range numbers {
		// Have to loop like this to prevent copying
		for _, grid := range grids {
			grid.Mark(n)
			if grid.IsWinning() {
				score := grid.GetScore()
				fmt.Printf("Score on winning grid: %d\n", n*score)
				return
			}
		}
	}
}
