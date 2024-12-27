package solutions

import (
	"fmt"
	"log"
)

//go:generate go run gen.go
var days map[int][]func(string)

func Execute(day int, input string) {
	if funcs, ok := days[day]; ok {
		for part, f := range funcs {
			fmt.Printf("Day %v, part %v: ", day+1, part+1)
			f(input)
		}
	} else {
		log.Fatalf("error: no such day: %v", day)
	}
}
