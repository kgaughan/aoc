package day10

import "testing"

type fixture struct {
	test     string
	expected string
}

func TestLookAndSay(t *testing.T) {
	tests := []fixture{
		{"1", "11"},
		{"11", "21"},
		{"21", "1211"},
		{"1211", "111221"},
		{"111221", "312211"},
	}

	for _, test := range tests {
		if result := LookAndSay(test.test); result != test.expected {
			t.Errorf(
				"For %q, expected %q, got %q",
				test.test, test.expected, result)
		}
	}
}
