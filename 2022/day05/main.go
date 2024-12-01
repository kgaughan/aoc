package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type stack struct {
	stack []rune
}

func (s *stack) pushUnder(r rune) {
	s.stack = append([]rune{r}, s.stack...)
}

func (s stack) top() rune {
	return s.stack[len(s.stack)-1]
}

func (s *stack) moveTo(other *stack, n int) bool {
	if len(s.stack) < n {
		return false
	}
	for i := 0; i < n; i++ {
		other.stack = append(other.stack, s.stack[len(s.stack)-i-1])
	}
	s.stack = s.stack[:len(s.stack)-n]
	return true
}

func (s *stack) moveBatch(other *stack, n int) bool {
	if len(s.stack) < n {
		return false
	}
	other.stack = append(other.stack, s.stack[len(s.stack)-n:]...)
	s.stack = s.stack[:len(s.stack)-n]
	return true
}

type rule struct {
	n, from, to int
}

func printTop(stacks []stack) {
	for i := 0; i < len(stacks); i++ {
		fmt.Printf("%c", stacks[i].top())
	}
}

func part1(stacks []stack, rules []rule) {
	for _, rule := range rules {
		stacks[rule.from-1].moveTo(&stacks[rule.to-1], rule.n)
	}
	fmt.Print("Part 1: ")
	printTop(stacks)
	fmt.Print("\n")
}

func part2(stacks []stack, rules []rule) {
	for _, rule := range rules {
		stacks[rule.from-1].moveBatch(&stacks[rule.to-1], rule.n)
	}
	fmt.Print("Part 2: ")
	printTop(stacks)
	fmt.Print("\n")
}

var (
	doPart1 = flag.Bool("1", false, "Do part 1")
	doPart2 = flag.Bool("2", false, "Do part 2")
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	stacks := make([]stack, 0, 10)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		// Ignore the last line
		if !strings.ContainsRune(line, '[') {
			continue
		}
		// Parsing this input is annoying
		i := 0
		runes := []rune(line)
		for iRune := 1; iRune < len(runes); iRune += 4 {
			if len(stacks) < i+1 {
				stacks = append(stacks, stack{})
			}
			if runes[iRune] != ' ' {
				stacks[i].pushUnder(rune(runes[iRune]))
			}
			i++
		}
	}

	// As we've used the scanner for the bits above, we need to continue using
	// it, or we'll skip data that it has buffered.
	rules := make([]rule, 0, 10)
	for scanner.Scan() {
		line := scanner.Text()
		var rule rule
		_, err := fmt.Sscanf(line, "move %d from %d to %d\n", &rule.n, &rule.from, &rule.to)
		if err == io.EOF {
			break
		}
		rules = append(rules, rule)
	}

	// The only reason I'm using flags here is that there's an information
	// leakage bug between parts one and two, and this was easier the fixing
	// it.
	flag.Parse()
	if *doPart1 {
		part1(stacks, rules)
	} else if *doPart2 {
		part2(stacks, rules)
	} else {
		flag.Usage()
	}
}
