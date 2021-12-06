package bingo

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func ReadNumbers(r *bufio.Reader) ([]int, error) {
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

func ReadGrids(r *bufio.Reader) ([]*Grid, error) {
	grids := make([]*Grid, 0, 100)
	for {
		// Skip empty line
		r.ReadString('\n')

		grid := &Grid{}
		if err := grid.Read(r); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		grids = append(grids, grid)
	}
	return grids, nil
}
