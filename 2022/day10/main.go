package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Machine struct {
	x        int
	cycle    int
	callback func(int, int)
}

func NewMachine(callback func(int, int)) *Machine {
	return &Machine{
		x:        1,
		cycle:    0,
		callback: callback,
	}
}

func (m *Machine) Execute(instruction string, operand int) {
	countdown := 0
	switch instruction {
	case "noop":
		countdown = 1
	case "addx":
		countdown = 2
	}

	for ; countdown > 0; countdown-- {
		m.cycle++
		m.callback(m.x, m.cycle)
		if countdown == 1 {
			switch instruction {
			case "addx":
				m.x += operand
			}
		}
	}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	signalSum := 0

	machine := NewMachine(func(x, cycle int) {
		if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
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
}
