package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func readInput() []byte {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func scan(data []byte, markerLen int) int {
	// A start of marker is n bytes that are all different.
	for i := 0; i < len(data)-markerLen; i++ {
		found := true
	scan:
		for j := 0; j < markerLen; j++ {
			for k := j + 1; k < markerLen; k++ {
				if data[i+j] == data[i+k] {
					found = false
					break scan
				}
			}
		}
		if found {
			return i + markerLen
		}
	}
	return -1
}

func main() {
	data := readInput()
	fmt.Printf("Start of packet found at %d\n", scan(data, 4))
	fmt.Printf("Start of message found at %d\n", scan(data, 14))
}
