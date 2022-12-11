package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type CPU struct {
	x        int
	cycle    int
	callback func(int, int)
	screen   *strings.Builder
}

func NewCPU(callback func(int, int)) *CPU {
	return &CPU{
		x:        1,
		cycle:    0,
		callback: callback,
		screen:   &strings.Builder{},
	}
}

func (c *CPU) Execute(instruction string, operand int) {
	countdown := 0
	switch instruction {
	case "noop":
		countdown = 1
	case "addx":
		countdown = 2
	}

	for ; countdown > 0; countdown-- {
		// Render the CRT
		gun := (c.cycle % screenWidth)
		if gun >= c.x-1 && gun <= c.x+1 {
			c.screen.WriteRune('#')
		} else {
			c.screen.WriteRune('.')
		}

		c.cycle++
		c.callback(c.x, c.cycle)

		// Run the instruction on the last cycle
		if countdown == 1 {
			switch instruction {
			case "addx":
				c.x += operand
			}
		}
	}
}

const screenWidth = 40

func (c CPU) String() string {
	screen := []rune(c.screen.String())
	result := &strings.Builder{}
	for i := 0; i < len(screen); i += screenWidth {
		result.WriteString(string(screen[i : i+screenWidth]))
		result.WriteRune('\n')
	}
	return result.String()
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	signalSum := 0

	machine := NewCPU(func(x, cycle int) {
		if cycle%screenWidth == 20 {
			fmt.Printf("X: %v; Cycle: %v\n", x, cycle)
			signalSum += x * cycle
		}
	})

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		instruction := scanner.Text()
		operand := 0
		if instruction == "addx" {
			if !scanner.Scan() {
				break
			}
			operand, err = strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}
		}
		machine.Execute(instruction, operand)
	}

	fmt.Printf("Part 1: %v\n", signalSum)
	fmt.Printf("Part 2:\n%v", machine)
}
