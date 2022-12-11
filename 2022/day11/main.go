package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type monkey struct {
	items       []int
	operator    rune
	operand     string
	divisor     int
	passToTrue  int
	passToFalse int
}

const monkeyPattern = `Monkey \d+:
  Starting items: ([0-9, ]+)
  Operation: new = old ([+*\-/]) (\d+|old)
  Test: divisible by (\d+)
    If true: throw to monkey (\d+)
    If false: throw to monkey (\d+)
`

func parseMonkey(matches []string) (*monkey, error) {
	var err error

	monkey := &monkey{}

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
	monkeys := make([]*monkey, 0, len(matches))
	for _, match := range matches {
		if monkey, err := parseMonkey(match); err != nil {
			log.Fatal(err)
		} else {
			monkeys = append(monkeys, monkey)
			fmt.Printf("%+v\n", monkey)
		}
	}
}
