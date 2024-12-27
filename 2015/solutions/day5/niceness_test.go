package day5

import "testing"

type fixture struct {
	s        string
	expected bool
}

func TestIsNice1(t *testing.T) {
	tests := []fixture{
		{"ugknbfddgicrmopn", true},
		{"aaa", true},
		{"jchzalrnumimnmhp", false},
		{"haegwjzuvuyypxyu", false},
		{"dvszwmarrgswjxmb", false},
	}

	driver(t, tests, "IsNice1", isNice1)
}

func TestIsNice2(t *testing.T) {
	tests := []fixture{
		{"xyxy", true},
		{"aaa", false},
		{"qjhvhtzxzqqjkmpb", true},
		{"xxyxx", true},
		{"uurcxstgmygtbstg", false},
		{"ieodomkazucvgmuy", false},
	}

	driver(t, tests, "IsNice2", isNice2)
}

func driver(t *testing.T, tests []fixture, fnName string, fn func(s string) bool) {
	for _, test := range tests {
		if answer := fn(test.s); answer != test.expected {
			t.Errorf(
				"%v(%q): got %v, expected %v",
				fnName, test.s, answer, test.expected)
		}
	}
}
