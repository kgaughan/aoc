package main

import (
	"aoc/day8"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	doQuote := flag.Bool("quote", false, "Quote the strings rather than unquoting them.")
	flag.Parse()

	var source io.Reader
	if flag.Arg(0) == "" {
		source = os.Stdin
	} else if file, err := os.Open(flag.Arg(0)); err != nil {
		log.Fatal(err)
	} else {
		defer file.Close()
		source = file
	}

	var conv func(string) string
	if *doQuote {
		conv = strconv.Quote
	} else {
		conv = func(s string) string {
			unquoted, _ := strconv.Unquote(s)
			return unquoted
		}
	}

	originalLength := 0
	convertedLength := 0
	day8.ScanLines(source, func(original string) {
		originalLength += len(original)
		convertedLength += len(conv(original))
	})

	difference := originalLength - convertedLength
	if originalLength < convertedLength {
		difference = -difference
	}
	fmt.Printf("Difference is %v\n", difference)
}
