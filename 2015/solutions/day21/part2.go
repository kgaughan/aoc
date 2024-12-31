package day21

import "fmt"

func Part2(input string) {
	bossHP, bossDamage, bossArmour := parse(input)
	highestCost := 0
	generateInventory(bossHP, bossDamage, bossArmour, func(playerWon bool, cost int) {
		if !playerWon && cost > highestCost {
			highestCost = cost
		}
	})
	fmt.Printf("%v\n", highestCost)
}
