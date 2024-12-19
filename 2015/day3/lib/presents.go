package lib

type point struct {
	x int
	y int
}

func Deliver(route string, n int) int {
	houses := make(map[point]int)
	coords := make([]point, n)

	// Visit all the houses.
	houses[point{0, 0}] = n
	for i, ch := range route {
		j := i % n
		switch ch {
		case 'v':
			coords[j].y -= 1
		case '^':
			coords[j].y += 1
		case '<':
			coords[j].x -= 1
		case '>':
			coords[j].x += 1
		}
		houses[coords[j]]++
	}

	// When we're done, the number of entries in the map will be the number of
	// houses visited...
	return len(houses)
}
