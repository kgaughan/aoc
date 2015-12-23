package day6

import "testing"

func assertLightCount(t *testing.T, lights *Lights, expected uint) {
	if n := lights.Count(); n != expected {
		t.Errorf("Expected %v lights, but counted %v", expected, n)
	}
}

func assertAddPoint(t *testing.T, lights *Lights, point Coord, expected uint) {
	lights.TurnOn(point, point)
	assertLightCount(t, lights, expected)
}

func TestLights(t *testing.T) {
	lights := Lights{}

	assertLightCount(t, &lights, 0)

	// Add initial dot.
	assertAddPoint(t, &lights, Coord{0, 0}, 1)

	// Setting it again shouldn't change the counter.
	assertAddPoint(t, &lights, Coord{0, 0}, 1)

	// Add three additional points.
	assertAddPoint(t, &lights, Coord{999, 0}, 2)
	assertAddPoint(t, &lights, Coord{0, 999}, 3)
	assertAddPoint(t, &lights, Coord{999, 999}, 4)

	// Flip everything.
	lights.Toggle(Coord{0,0}, Coord{999,999})
	assertLightCount(t, &lights, 999996)

	// Turn everything off.
	lights.TurnOff(Coord{0,0}, Coord{999,999})
	assertLightCount(t, &lights, 0)
}
