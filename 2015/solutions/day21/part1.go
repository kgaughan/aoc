package day21

import (
	"fmt"
	"math"
)

func Part1(input string) {
	bossHP, bossDamage, bossArmour := parse(input)
	lowestCost := math.MaxInt32
	generateInventory(bossHP, bossDamage, bossArmour, func(playerWon bool, cost int) {
		if playerWon && cost < lowestCost {
			lowestCost = cost
		}
	})
	fmt.Printf("%v\n", lowestCost)
}
