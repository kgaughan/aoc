package day6

type Coord struct {
	X, Y int
}

const max = 1000

type Lights struct {
	lights [max][max]uint8
}

func apply(from, to Coord, fn func(int, int)) {
	for x := from.X; x <= to.X; x++ {
		for y := from.Y; y <= to.Y; y++ {
			fn(x, y)
		}
	}
}

func (l *Lights) TurnOn(from, to Coord) {
	apply(from, to, func(x, y int) {
		l.lights[x][y] = 1
	})
}

func (l *Lights) TurnOff(from, to Coord) {
	apply(from, to, func(x, y int) {
		l.lights[x][y] = 0
	})
}

func (l *Lights) Toggle(from, to Coord) {
	apply(from, to, func(x, y int) {
		l.lights[x][y] = 1 - l.lights[x][y]
	})
}

func (l *Lights) Count() uint {
	var n uint

	n = 0
	for x := 0; x < max; x++ {
		for y := 0; y < max; y++ {
			n += uint(l.lights[x][y])
		}
	}
	return n
}
