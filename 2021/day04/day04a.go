package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/kgaughan/aoc/2021/day04/bingo"
)

func readNumbers(r *bufio.Reader) ([]int, error) {
	line, err := r.ReadString('\n')
	if err != nil {
		return nil, err
	}
	numbers := strings.Split(strings.TrimSpace(line), ",")
	parsedNumbers := make([]int, len(numbers))
	for i, n := range numbers {
		parsed, err := strconv.ParseInt(n, 10, 64)
		if err != nil {
			return nil, err
		}
		parsedNumbers[i] = int(parsed)
	}
	return parsedNumbers, nil
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	numbers, err := readNumbers(reader)
	if err != nil {
		log.Fatal(err)
	}

	grids := make([]bingo.Grid, 0, 100)
	for {
		// Skip empty line
		reader.ReadString('\n')

		grid := bingo.Grid{}
		if err := grid.Read(reader); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		grids = append(grids, grid)
	}

	for _, n := range numbers {
		// Have to loop like this to prevent copying
		for i := 0; i < len(grids); i++ {
			grids[i].Mark(n)
			if grids[i].IsWinning() {
				score := grids[i].GetScore()
				fmt.Printf("Score on winning grid: %d\n", n*score)
				return
			}
		}
	}
}
