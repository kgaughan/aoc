package day21

import (
	"strconv"
	"strings"

	"github.com/kgaughan/aoc/2015/helpers"
)

const maxHP = 100

type item struct {
	name                 string
	cost, damage, armour int
}

var dummy = item{}

var weapons = []item{
	{name: "Dagger", cost: 8, damage: 4},
	{name: "Shortsword", cost: 10, damage: 5},
	{name: "Warhammer", cost: 25, damage: 6},
	{name: "Longsword", cost: 40, damage: 7},
	{name: "Greataxe", cost: 74, damage: 8},
}

var armour = []item{
	dummy,
	{name: "Leather", cost: 13, armour: 1},
	{name: "Chainmail", cost: 31, armour: 2},
	{name: "Splintmail", cost: 53, armour: 3},
	{name: "Bandedmail", cost: 75, armour: 4},
	{name: "Platemail", cost: 102, armour: 5},
}

var rings = []item{
	dummy,
	{name: "Damage +1", cost: 25, damage: 1},
	{name: "Damage +2", cost: 50, damage: 2},
	{name: "Damage +3", cost: 100, damage: 3},
	{name: "Defense +1", cost: 20, armour: 1},
	{name: "Defense +2", cost: 40, armour: 2},
	{name: "Defense +3", cost: 80, armour: 3},
}

func parse(input string) (int, int, int) {
	bossHP := 0
	bossDamage := 0
	bossArmour := 0
	helpers.ScanLines(strings.NewReader(input), func(s string) error {
		parts := strings.SplitN(s, ": ", 2)
		value, err := strconv.Atoi(parts[1])
		if err != nil {
			return err
		}
		switch parts[0] {
		case "Hit Points":
			bossHP = value
		case "Damage":
			bossDamage = value
		case "Armor":
			bossArmour = value
		}
		return nil
	})

	return bossHP, bossDamage, bossArmour
}

func generateInventory(bossHP, bossDamage, bossArmour int, evaluate func(playerWon bool, cost int)) {
	gear := make([]item, 4) // A weapon, armour, and two rings
	for _, weapon := range weapons {
		gear[0] = weapon
		for _, armourItem := range armour {
			gear[1] = armourItem
			for i, ring := range rings {
				gear[2] = ring
				simulate(bossHP, bossDamage, bossArmour, gear, evaluate)
				for j := i + 1; j < len(rings); j++ {
					gear[3] = rings[j]
					simulate(bossHP, bossDamage, bossArmour, gear, evaluate)
				}
				gear[3] = dummy
			}
		}
	}
}

func simulate(bossHP, bossDamage, bossArmour int, gear []item, fn func(playerWon bool, cost int)) {
	playerDMG := 0
	playerDEF := 0
	cost := 0
	for _, item := range gear {
		playerDMG += item.damage
		playerDEF += item.armour
		cost += item.cost
	}
	playerWins := bossHP/max(playerDMG-bossArmour, 1) <= maxHP/max(bossDamage-playerDEF, 1)
	fn(playerWins, cost)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
