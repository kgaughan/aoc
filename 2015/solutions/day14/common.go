package day14

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

const raceDuration = 2503

type stats struct{ speed, duration, rest int }

func parseEntry(source io.Reader) (string, stats, error) {
	var name string
	var speed int
	var duration int
	var rest int
	_, err := fmt.Fscanf(
		source,
		"%s can fly %d km/s for %d seconds, but then must rest for %d seconds.\n",
		&name, &speed, &duration, &rest)
	return name, stats{speed: speed, duration: duration, rest: rest}, err
}

func parse(input string) map[string]stats {
	competitors := make(map[string]stats, 10)
	reader := strings.NewReader(input)
	for {
		if name, stats, err := parseEntry(reader); err == nil {
			competitors[name] = stats
		} else if errors.Is(err, io.EOF) {
			break
		} else {
			panic(err)
		}
	}
	return competitors
}
