package main

import (
	"bufio"
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
	for scanner.Scan() {
		line := scanner.Text()
		var n, from, to int
		_, err := fmt.Sscanf(line, "move %d from %d to %d\n", &n, &from, &to)
		if err == io.EOF {
			break
		}
		stacks[from-1].moveTo(&stacks[to-1], n)
	}
	for i := 0; i < len(stacks); i++ {
		fmt.Printf("%c", stacks[i].top())
	}
	fmt.Print("\n")
}
