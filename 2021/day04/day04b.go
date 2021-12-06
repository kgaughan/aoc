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

	var lastWinningGrid *bingo.Grid
	eliminated := make(map[int]bool, len(grids))
	lastWinningNumber := 0
	for _, n := range numbers {
		for i, grid := range grids {
			if _, exists := eliminated[i]; exists {
				continue
			}
			grid.Mark(n)
			if grid.IsWinning() {
				eliminated[i] = true
				lastWinningGrid = grid
				lastWinningNumber = n
			}
		}
	}

	fmt.Printf("Score on winning grid: %d\n", lastWinningGrid.GetScore()*lastWinningNumber)
}
