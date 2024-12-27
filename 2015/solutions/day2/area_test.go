package day2

import "testing"

type fixture struct {
	d        dimensions
	expected int
}

func TestArea(t *testing.T) {
	tests := []fixture{
		{dimensions{2, 3, 4}, 58},
		{dimensions{1, 1, 10}, 43},
	}
	driver(t, tests, "area", func(d dimensions) int { return d.area() })
}

func TestRibbon(t *testing.T) {
	tests := []fixture{
		{dimensions{2, 3, 4}, 34},
		{dimensions{1, 1, 10}, 14},
	}
	driver(t, tests, "ribbon", func(d dimensions) int { return d.ribbon() })
}

func driver(t *testing.T, tests []fixture, name string, f func(d dimensions) int) {
	for _, test := range tests {
		if answer := f(test.d); answer != test.expected {
			t.Errorf(
				"%v(%v, %v, %v): got %v, expected %v",
				name, test.d.l, test.d.w, test.d.h,
				answer, test.expected)
		}
	}
}
