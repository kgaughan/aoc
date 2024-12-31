package day15

import "fmt"

func Part1(input string) {
	ingredients := parse(input)
	cookie := NewRecipe(ingredients)

	// From looking stuff up later, I believe this is an implementation of hill
	// climbing, but I honestly don't know.
	for {
		adjusted := false
		for i := 0; i < len(ingredients); i++ {
			if cookie.amounts[i] > 0 {
				for j := 0; j < len(ingredients); j++ {
					if i != j {
						trial := cookie.swap(i, j)
						if trial.score(ingredients) > cookie.score(ingredients) {
							adjusted = true
							cookie = trial
						}
					}
				}
			}
		}
		if !adjusted {
			break
		}
	}

	fmt.Printf("%v\n", cookie.score(ingredients))
}
