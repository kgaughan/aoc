package day15

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

const totalTeaspoons = 100

type ingredient struct {
	name                                             string
	capacity, durability, flavour, texture, calories int
}

func (i ingredient) getScores(amt int) (int, int, int, int) {
	return amt * i.capacity, amt * i.durability, amt * i.flavour, amt * i.texture
}

func parseEntry(source io.Reader) (ingredient, error) {
	var e ingredient
	_, err := fmt.Fscanf(
		source,
		"%s capacity %d, durability %d, flavor %d, texture %d, calories %d\n",
		&e.name, &e.capacity, &e.durability, &e.flavour, &e.texture, &e.calories)
	if err == nil {
		e.name = strings.TrimRight(e.name, ":")
	}
	return e, err
}

func parse(input string) []ingredient {
	ingredients := make([]ingredient, 0, 10)
	reader := strings.NewReader(input)
	for {
		if e, err := parseEntry(reader); err == nil {
			ingredients = append(ingredients, e)
		} else if errors.Is(err, io.EOF) {
			break
		} else {
			panic(err)
		}
	}
	return ingredients
}

type Recipe struct {
	amounts []int
}

func NewRecipe(ingredients []ingredient) *Recipe {
	// Let's assume even cookie are a resonable starting point
	cookie := &Recipe{
		amounts: make([]int, len(ingredients)),
	}
	for i := 0; i < len(cookie.amounts); i++ {
		cookie.amounts[i] = totalTeaspoons / len(cookie.amounts)
	}
	// A fudge in case the number of teaspoons isn't evenly divisible
	cookie.amounts[0] += totalTeaspoons % len(cookie.amounts)
	return cookie
}

func (r Recipe) swap(i, j int) *Recipe {
	result := &Recipe{
		amounts: make([]int, len(r.amounts)),
	}
	copy(result.amounts, r.amounts)
	result.amounts[i]--
	result.amounts[j]++
	return result
}

func (r Recipe) calories(ingredients []ingredient) int {
	calories := 0
	for i, amount := range r.amounts {
		calories += ingredients[i].calories * amount
	}
	return calories
}

func (r Recipe) score(ingredients []ingredient) int {
	capacity := 0
	durability := 0
	flavour := 0
	texture := 0
	for i, amount := range r.amounts {
		c, d, f, t := ingredients[i].getScores(amount)
		capacity += c
		durability += d
		flavour += f
		texture += t
	}
	if capacity < 0 || durability < 0 || flavour < 0 || texture < 0 {
		return 0
	}
	return capacity * durability * flavour * texture
}
