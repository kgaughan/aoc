package day17

import (
	"strconv"
	"strings"

	"github.com/kgaughan/aoc/2015/helpers"
)

const eggnog = 150

func readInput(input string) []int {
	containers := make([]int, 0, 20)
	helpers.ScanLines(strings.NewReader(input), func(s string) {
		if value, err := strconv.Atoi(s); err != nil {
			panic(err)
		} else {
			containers = append(containers, value)
		}
	})
	return containers
}
