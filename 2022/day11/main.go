package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	items       []int
	operator    rune
	operand     string
	divisor     int
	passToTrue  int
	passToFalse int
	inspections int
}

const monkeyPattern = `Monkey \d+:
  Starting items: ([0-9, ]+)
  Operation: new = old ([+*\-/]) (\d+|old)
  Test: divisible by (\d+)
    If true: throw to monkey (\d+)
    If false: throw to monkey (\d+)
`

func extractMonkey(matches []string) (*Monkey, error) {
	var err error

	monkey := &Monkey{}

	splitItems := strings.Split(matches[1], ",")
	monkey.items = make([]int, 0, len(splitItems))
	for _, item := range splitItems {
		n, err := strconv.Atoi(strings.TrimSpace(item))
		if err != nil {
			return nil, err
		}
		monkey.items = append(monkey.items, n)
	}

	monkey.operator = []rune(matches[2])[0]
	monkey.operand = matches[3]

	if monkey.divisor, err = strconv.Atoi(matches[4]); err != nil {
		return nil, err
	}

	if monkey.passToTrue, err = strconv.Atoi(matches[5]); err != nil {
		return nil, err
	}

	if monkey.passToFalse, err = strconv.Atoi(matches[6]); err != nil {
		return nil, err
	}

	return monkey, nil
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(monkeyPattern)
	matches := re.FindAllStringSubmatch(string(data), -1)
	monkeys := make([]*Monkey, 0, len(matches))
	for _, match := range matches {
		if monkey, err := extractMonkey(match); err != nil {
			log.Fatal(err)
		} else {
			monkeys = append(monkeys, monkey)
		}
	}

	for i := 0; i < 20; i++ {
		log.Printf("Round %d", i+1)
		for j, monke := range monkeys {
			log.Printf("  Monkey %d:", j)
			for _, worry := range monke.items {
				monke.inspections++
				log.Printf("    Inspecting item with worry level of %d", worry)
				var operand int
				if monke.operand == "old" {
					operand = worry
				} else if n, err := strconv.Atoi(monke.operand); err == nil {
					operand = n
				} else {
					log.Fatalf("Bad operand: %q", monke.operand)
				}
				switch monke.operator {
				case '+':
					worry += operand
				case '-':
					worry -= operand
				case '*':
					worry *= operand
				case '/':
					worry /= operand
				}
				log.Printf("      Applying %c %v; worry level is now %d", monke.operator, monke.operand, worry)
				worry /= 3
				log.Printf("      Inspecting after reducing worry level to %d", worry)
				var iReceiver int
				if worry%monke.divisor == 0 {
					iReceiver = monke.passToTrue
				} else {
					iReceiver = monke.passToFalse
				}
				log.Printf("      Passing to %d", iReceiver)
				monkeys[iReceiver].items = append(monkeys[iReceiver].items, worry)
			}
			monke.items = monke.items[:0]
		}
	}

	inspections := make([]int, len(monkeys))
	for i, monkey := range monkeys {
		inspections[i] = monkey.inspections
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspections)))

	fmt.Printf("Part 1: %v monkey business\n", inspections[0]*inspections[1])
}
