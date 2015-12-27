package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"aoc/day8"
)

func main() {
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

	originalLength := 0
	unquotedLength := 0
	day8.ParseStrings(source, func(original, unquoted string) {
		originalLength += len(original)
		unquotedLength += len(unquoted)
	})
	fmt.Printf("Difference is %v\n", originalLength - unquotedLength)
}
