package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func csvSplit(line string) []string {
	return strings.Split(strings.TrimSpace(line), ",")
}

func parseFloats(in []string) ([]float64, error) {
	out := make([]float64, len(in))
	for i, n := range in {
		parsed, err := strconv.ParseFloat(n, 64)
		if err != nil {
			return nil, err
		}
		out[i] = parsed
	}
	return out, nil
}

func harmonicMean(in []float64) float64 {
	var sum float64 = 0.0
	for _, n := range in {
		if n != 0.0 {
			sum += 1 / n
		}
	}
	return float64(len(in)) / sum
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	numbers, err := parseFloats(csvSplit(line))
	if err != nil {
		log.Fatal(err)
	}

	/*
		// Calculate the harmonic mean. Why the harmonic mean? Well, because
		// that's the method used to calculate the resistance in parallel,
		// and this looks a lot like that. The harmonic mean would be the
		// position all the crabs need to move to in order to minimise fuel
		// consumption amongst all of them.
		pos := math.Round(harmonicMean(numbers))

	*/
	sort.Float64s(numbers)
	pos := numbers[len(numbers)/2]

	fuel := 0.0
	for _, n := range numbers {
		fuel += math.Abs(n - pos)
	}
	fmt.Printf("Fuel use: %v\n", fuel)
}
