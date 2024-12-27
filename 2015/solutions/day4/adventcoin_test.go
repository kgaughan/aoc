package day4

import "testing"

type fixture struct {
	prefix   string
	expected int
}

func TestAdventCoin(t *testing.T) {
	tests := []fixture{
		{"abcdef", 609043},
		{"pqrstuv", 1048970},

		// Supplied prefix and its answer.
		{"ckczppom", 117946},
	}

	for _, test := range tests {
		if answer := adventCoin(test.prefix, 5, test.expected-1, test.expected+1); answer != test.expected {
			t.Errorf("adventCoin(%q): got %v; expected %v", test.prefix, answer, test.expected)
		}
	}
}
