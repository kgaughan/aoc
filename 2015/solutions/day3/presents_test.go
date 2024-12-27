package day3

import "testing"

type fixture struct {
	route    string
	expected int
}

func TestDeliver1(t *testing.T) {
	tests := []fixture{
		{"", 1},
		{">", 2},
		{">v<^", 4},
		{"^v^v^v^v^v", 2},
		{">>>>", 5},
		{"<<<<", 5},
		{"vvvv", 5},
		{"^^^^", 5},
	}
	driver(t, tests, 1)
}

func TestDeliver2(t *testing.T) {
	tests := []fixture{
		{"", 1},
		{">", 2},
		{">v<^", 3},
		{"^v^v^v^v^v", 11},
		{">>>>", 3},
		{"<<<<", 3},
		{"vvvv", 3},
		{"^^^^", 3},
		{"^v", 3},
	}
	driver(t, tests, 2)
}

func driver(t *testing.T, tests []fixture, n int) {
	for _, test := range tests {
		if answer := deliver(test.route, n); answer != test.expected {
			t.Errorf(
				"Deliver(%q, %v): got %v; expected %v",
				test.route, n,
				answer, test.expected)
		}
	}
}
