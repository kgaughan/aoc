package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const DAYS = 256 // this is for part 2; use 80 for part 1

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	lanternfish := make([]int, 9)

	reader := bufio.NewReader(f)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	numbers := strings.Split(strings.TrimSpace(line), ",")
	for _, n := range numbers {
		idx, err := strconv.ParseInt(n, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		lanternfish[idx]++
	}

	for i := 0; i < 256; i++ {
		toDouble := lanternfish[0]
		for j := 1; j < len(lanternfish); j++ {
			lanternfish[j-1] = lanternfish[j]
		}
		// Cycle the parents
		lanternfish[6] += toDouble
		// Add the offspring
		lanternfish[8] = toDouble
	}

	sum := 0
	for _, n := range lanternfish {
		sum += n
	}

	fmt.Printf("Lanternfish after %d days: %d\n", DAYS, sum)
}
