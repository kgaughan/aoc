package day20

import (
	"fmt"
	"math"
	"strconv"
)

func Part1(input string) {
	target, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}

	houseNumber := 1
	for {
		received := 0
		max := int(math.Sqrt(float64(houseNumber)))
		for i := 1; i <= max; i++ {
			if houseNumber%i == 0 {
				received += i
				if houseNumber/i != i {
					received += houseNumber / i
				}
			}
		}

		if received*10 >= target {
			break
		}
		houseNumber++
	}
	fmt.Printf("%v\n", houseNumber)
}
