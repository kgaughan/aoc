package day14

import "fmt"

type state struct {
	duration, distance, score int
	resting                   bool
}

func Part2(input string) {
	competitors := parse(input)

	states := make(map[string]*state, len(competitors))

	for name, s := range competitors {
		states[name] = &state{duration: s.duration, distance: 0, score: 0, resting: false}
	}

	for remaining := raceDuration; remaining > 0; remaining-- {
		furthest := 0

		for name, s := range states {
			if !s.resting {
				s.distance += competitors[name].speed
			}
			if s.distance > furthest {
				furthest = s.distance
			}
			s.duration--
			if s.duration == 0 {
				if s.resting {
					s.duration = competitors[name].duration
				} else {
					s.duration = competitors[name].rest
				}
				s.resting = !s.resting
			}
		}

		for _, s := range states {
			if s.distance == furthest {
				s.score++
			}
		}
	}

	maxScore := 0
	for _, s := range states {
		if s.score > maxScore {
			maxScore = s.score
		}
	}

	fmt.Printf("Best: %v\n", maxScore)
}
