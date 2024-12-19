package lib

import "testing"

func TestCountFloors(t *testing.T) {
	tests := []struct {
		directions string
		floors     int
	}{
		{"(())", 0},
		{"()()", 0},
		{"(((", 3},
		{"(()(()(", 3},
		{"))(((((", 3},
		{"())", -1},
		{"))(", -1},
		{")))", -3},
		{")())())", -3},
		{"", 0},
		{"(", 1},
		{")", -1},
	}

	for _, test := range tests {
		if answer := CountFloors(test.directions); test.floors != answer {
			t.Errorf("CountFloors: For %q, expected %v, got %v", test.directions, test.floors, answer)
		}
	}
}

func TestFindBasementInstruction(t *testing.T) {
	tests := []struct {
		directions string
		position   int
	}{
		{")", 1},
		{"()())", 5},
	}

	for _, test := range tests {
		if answer := FindBasementInstruction(test.directions); test.position != answer {
			t.Errorf("FindBasementInstruction: For %q, expected %v, got %v", test.directions, test.position, answer)
		}
	}
}
