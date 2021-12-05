package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var depth int64 = 0
	var distance int64 = 0
	var aim int64 = 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, " ", 2)
		amount, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		switch parts[0] {
		case "forward":
			distance += amount
		case "up":
			depth -= amount
		case "down":
			depth += amount
		}
	}
	fmt.Printf("Product: %v\n", depth*distance)
}
