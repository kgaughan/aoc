package day12

import (
	"encoding/json"
	"log"
)

func parse(input string) interface{} {
	var obj interface{}
	if err := json.Unmarshal([]byte(input), &obj); err != nil {
		log.Fatal(err)
	}
	return obj
}
