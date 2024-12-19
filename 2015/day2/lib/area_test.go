package lib

import "testing"

type fixture struct {
	l        int
	w        int
	h        int
	expected int
}

func TestArea(t *testing.T) {
	tests := []fixture{
		{2, 3, 4, 58},
		{1, 1, 10, 43},
	}
	driver(t, tests, "Area", Area)
}

func TestRibbon(t *testing.T) {
	tests := []fixture{
		{2, 3, 4, 34},
		{1, 1, 10, 14},
	}
	driver(t, tests, "Ribbon", Ribbon)
}

func driver(t *testing.T, tests []fixture, name string, f func(l, w, h int) int) {
	for _, test := range tests {
		if answer := f(test.l, test.w, test.h); answer != test.expected {
			t.Errorf(
				"%v(%v, %v, %v): got %v, expected %v",
				name, test.l, test.w, test.h,
				answer, test.expected)
		}
	}
}
