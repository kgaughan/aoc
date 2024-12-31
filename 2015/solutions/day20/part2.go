package day20

import (
	"fmt"
	"math"
	"strconv"
)

func Part2(input string) {
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
				if houseNumber/i <= 50 {
					received += i
				}
				if i <= 50 && houseNumber/i != i {
					received += houseNumber / i
				}
			}
		}

		if received*11 >= target {
			break
		}
		houseNumber++
	}
	fmt.Printf("%v\n", houseNumber)
}
