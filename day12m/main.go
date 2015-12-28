package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"aoc/day12"
)

type pair struct {
	from, to string
}

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

	doc, err := ioutil.ReadAll(source)
	if err != nil {
		log.Fatal(err)
	}

	var obj interface{}
	if err := json.Unmarshal(doc, &obj); err != nil {
		log.Fatal(err)
	}

	result := day12.AddNumbers(obj)
	fmt.Printf("Total is: %v\n", result)
}
