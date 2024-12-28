package day9

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

func parse(input string) (Edges, []string) {
	edges := Edges{}
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
		} else if errors.Is(err, io.EOF) {
			return nil
		} else {
			return err
		}
	}
}
