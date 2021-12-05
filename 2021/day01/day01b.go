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
	var previousWindow int64 = 0
	var previous1 int64 = 0
	var previous2 int64 = 0
	total := 0
	for n := 0; scanner.Scan(); n++ {
		line := scanner.Text()
		current, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		if n > 2 {
			newWindow := current + previous1 + previous2
			if previousWindow < newWindow {
				total++
				log.Println(">")
			} else {
				log.Println("<=")
			}
			previousWindow = newWindow
		}
		previous2 = previous1
		previous1 = current
	}
	fmt.Printf("Increases: %v\n", total)
}
