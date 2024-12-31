package day15

import (
	"fmt"
)

const maxCalories = 500

func Part2(input string) {
	ingredients := parse(input)
	cookie := NewRecipe(ingredients)

	// Avoid attempts that would blow our calorie budget.
	maxIterations := make([]int, len(cookie.amounts))
	for i, ingredient := range ingredients {
		maxIterations[i] = min(totalTeaspoons, maxCalories/ingredient.calories)
	}

	largest := trial(ingredients, cookie, maxIterations, 0, totalTeaspoons)
	fmt.Printf("%v\n", largest)
}

// I hate this, but all the other methods I tried didn't work as I'd hoped. I
// kept hitting local minima when I tried to do gradient descent. Brute force
// it is, then.
func trial(ingredients []ingredient, cookie *Recipe, maxIterations []int, depth, remaining int) int {
	if depth == len(ingredients)-1 || remaining == 0 {
		if remaining <= maxIterations[depth] {
			cookie.amounts[depth] = remaining
			if cookie.calories(ingredients) == maxCalories {
				return cookie.score(ingredients)
			}
		}
		return 0
	}
	largest := 0
	for i := 0; i <= min(maxIterations[depth], remaining); i++ {
		cookie.amounts[depth] = i
		score := trial(ingredients, cookie, maxIterations, depth+1, remaining-i)
		if score > largest {
			largest = score
		}
	}
	return largest
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
