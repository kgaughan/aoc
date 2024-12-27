package day12

import (
	"encoding/json"
	"fmt"
	"log"
)

type pair struct {
	from, to string
}

func Part1(input string) {
	var obj interface{}
	if err := json.Unmarshal([]byte(input), &obj); err != nil {
		log.Fatal(err)
	}

	result := AddNumbers(obj)
	fmt.Printf("Total is: %v\n", result)
}
