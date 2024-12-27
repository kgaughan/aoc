package day9

import (
	"fmt"
	"io"
	"strings"

	"github.com/kgaughan/aoc/2015/helpers/graph"
)

func parse(input string) (graph.Edges, []string) {
	edges := graph.Edges{}
	towns := make([]string, 0, 50)
	ScanLines(strings.NewReader(input), func(from, to string, distance int) {
		if !Contains(towns, from) {
			towns = append(towns, from)
		}
		if !Contains(towns, to) {
			towns = append(towns, to)
		}
		edges.Add(from, to, distance)
	})
	return edges, towns
}

func ScanLines(reader io.Reader, receive func(string, string, int)) error {
	var from, to string
	var distance int
	for {
		if _, err := fmt.Fscanf(reader, "%s to %s = %d\n", &from, &to, &distance); err == nil {
			receive(from, to, distance)
		} else if err == io.EOF {
			return nil
		} else {
			return err
		}
	}
}
