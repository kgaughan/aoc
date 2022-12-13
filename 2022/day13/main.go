package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
)

const (
	CONTINUE = iota
	FAIL
	IN_ORDER
)

func compare(left, right any) int {
	switch leftValue := left.(type) {
	case float64:
		switch rightValue := right.(type) {
		case float64:
			if leftValue > rightValue {
				return FAIL
			}
			if leftValue < rightValue {
				return IN_ORDER
			}
			return CONTINUE
		case []any:
			return compare([]any{leftValue}, rightValue)
		default:
			log.Fatalf("Unexpected type!")
		}
	case []any:
		switch rightValue := right.(type) {
		case float64:
			return compare(leftValue, []any{rightValue})
		case []any:
			for i, lv := range leftValue {
				if i == len(rightValue) {
					return FAIL
				}
				result := compare(lv, rightValue[i])
				if result == FAIL || result == IN_ORDER {
					return result
				}
			}
			if len(leftValue) == len(rightValue) {
				return CONTINUE
			}
			return IN_ORDER
		default:
			log.Fatalf("Unexpected type!")
		}
	default:
		log.Fatalf("Unexpected type!")
	}
	log.Fatal("Should be unreachable!")
	return FAIL
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	packets := make([]any, 0, 10)
	for scanner.Scan() {
		if len(scanner.Bytes()) == 0 {
			continue
		}
		var lst []any
		if err := json.Unmarshal(scanner.Bytes(), &lst); err != nil {
			log.Fatal(err)
		}
		packets = append(packets, lst)
	}

	indexSum := 0
	for i := 0; i < len(packets); i += 2 {
		if compare(packets[i], packets[i+1]) != FAIL {
			indexSum += (i / 2) + 1
		}
	}

	fmt.Printf("Part 1: %v\n", indexSum)

	dividerPackets := []any{
		[]any{[]any{2.0}},
		[]any{[]any{6.0}},
	}

	packets = append(packets, dividerPackets...)
	sort.Slice(packets, func(i, j int) bool {
		return compare(packets[i], packets[j]) != FAIL
	})
	decoderKey := 1
	for i, packet := range packets {
		for _, divider := range dividerPackets {
			if fmt.Sprint(packet) == fmt.Sprint(divider) {
				decoderKey *= i + 1
			}
		}
	}
	fmt.Printf("Part 2: %v\n", decoderKey)
}
