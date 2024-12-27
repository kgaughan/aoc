package day1

func toDirection(direction rune) int {
	switch direction {
	case '(':
		return 1
	case ')':
		return -1
	default:
		return 0
	}
}

func countFloors(directions string) int {
	floor := 0
	for _, c := range directions {
		floor += toDirection(c)
	}
	return floor
}

func findBasementInstruction(directions string) int {
	floor := 0
	for i, c := range directions {
		floor += toDirection(c)
		if floor == -1 {
			// Convert 0-index t o 1-index
			return i + 1
		}
	}
	return -1
}
