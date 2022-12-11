package main

import (
	"flag"
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

// Greatest common denominator
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Least common multiple: needed to prevent the numbers from getting out of hand.
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

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

var worryReduction = flag.Int("reduction", 3, "worry reduction")
var rounds = flag.Int("rounds", 20, "rounds to play")

func main() {
	flag.Parse()

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

	leastCommonMultiple := 1
	for _, monke := range monkeys {
		leastCommonMultiple = lcm(leastCommonMultiple, monke.divisor)
	}

	for i := 0; i < *rounds; i++ {
		for _, monke := range monkeys {
			for _, worry := range monke.items {
				monke.inspections++
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
				case '*':
					worry *= operand
				}
				worry /= *worryReduction
				// LCM adjustment to prevent overflows...
				worry %= leastCommonMultiple
				var iReceiver int
				if worry%monke.divisor == 0 {
					iReceiver = monke.passToTrue
				} else {
					iReceiver = monke.passToFalse
				}
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

	fmt.Printf("%v monkey business\n", inspections[0]*inspections[1])
}
