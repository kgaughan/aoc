package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	nLines := 0
	totals := make([]int, 0, 256)
	for scanner.Scan() {
		line := scanner.Text()
		for idx, ch := range line {
			if len(totals) <= idx {
				totals = append(totals, 0)
			}
			if ch == '1' {
				totals[idx] += 1
			}
		}
		nLines++
	}
	gamma := 0
	epsilon := 0
	for _, n := range totals {
		gamma <<= 1
		epsilon <<= 1
		if n > nLines/2 {
			gamma |= 1
		} else {
			epsilon |= 1
		}
	}

	fmt.Printf("Product: %v\n", gamma*epsilon)
}
