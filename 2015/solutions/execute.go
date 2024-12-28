package solutions

import (
	"fmt"
	"log"
	"time"
)

//go:generate go run gen.go
var days map[int][]func(string)

func Execute(day int, input string) {
	if funcs, ok := days[day]; ok {
		for part, f := range funcs {
			fmt.Printf("Day %v, part %v: ", day, part+1)
			func() {
				start := time.Now()
				defer func() {
					fmt.Printf("Duration: %v\n", time.Since(start))
				}()
				f(input)
			}()
		}
	} else {
		log.Fatalf("error: no such day: %v", day)
	}
}
