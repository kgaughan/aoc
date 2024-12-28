package day3

type point struct{ x, y int }

func deliver(route string, n int) int {
	houses := make(map[point]int)
	coords := make([]point, n)

	// Visit all the houses.
	houses[point{0, 0}] = n
	for i, ch := range route {
		j := i % n
		switch ch {
		case 'v':
			coords[j].y--
		case '^':
			coords[j].y++
		case '<':
			coords[j].x--
		case '>':
			coords[j].x++
		}
		houses[coords[j]]++
	}

	// When we're done, the number of entries in the map will be the number of
	// houses visited...
	return len(houses)
}
