package day1

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
		if answer := countFloors(test.directions); test.floors != answer {
			t.Errorf("countFloors: For %q, expected %v, got %v", test.directions, test.floors, answer)
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
		if answer := findBasementInstruction(test.directions); test.position != answer {
			t.Errorf("findBasementInstruction: For %q, expected %v, got %v", test.directions, test.position, answer)
		}
	}
}
