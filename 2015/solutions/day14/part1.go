package day14

import (
	"fmt"
)

func Part1(input string) {
	competitors := parse(input)

	best := 0
	for _, s := range competitors {
		// Lowball by assuming a flight and rest period
		stages := raceDuration / (s.duration + s.rest)
		distance := s.speed * s.duration * stages
		// Compensate if the reindeer could've flown a bit more
		remainder := raceDuration - (s.rest+s.duration)*stages
		if remainder < s.duration {
			distance += remainder * s.speed
		} else {
			distance += s.duration * s.speed
		}

		if distance > best {
			best = distance
		}
	}

	fmt.Printf("Best: %v\n", best)
}
