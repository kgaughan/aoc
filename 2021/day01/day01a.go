package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	hasPrevious := false
	var previous int64 = 0
	n := 0
	for scanner.Scan() {
		line := scanner.Text()
		current, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		if hasPrevious {
			if previous < current {
				n++
				log.Println(">")
			} else {
				log.Println("<=")
			}
		}
		previous = current
		hasPrevious = true
	}
	fmt.Printf("Increases: %v\n", n)
}
