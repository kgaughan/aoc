package day5

import "testing"

type fixture struct {
	s        string
	expected bool
}

func TestIsNice(t *testing.T) {
	tests := []fixture{
		{"ugknbfddgicrmopn", true},
		{"aaa", true},
		{"jchzalrnumimnmhp", false},
		{"haegwjzuvuyypxyu", false},
		{"dvszwmarrgswjxmb", false},
	}

	for _, test := range tests {
		if answer := IsNice(test.s); answer != test.expected {
			t.Errorf(
				"IsNice(%q): got %v, expected %v",
				test.s, answer, test.expected)
		}
	}
}
