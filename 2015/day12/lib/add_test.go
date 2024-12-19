package lib

import (
	"encoding/json"
	"testing"
)

type fixture struct {
	doc      string
	expected float64
}

func TestAddNumbers(t *testing.T) {
	tests := []fixture{
		// Part 1 tests.
		{`[1,2,3]`, 6},
		{`{"a":2,"b":4}`, 6},
		{`[[[3]]]`, 3},
		{`{"a":{"b":4},"c":-1}`, 3},
		{`{"a":[-1,1]}`, 0},
		{`[-1,{"a":1}]`, 0},
		{`{}`, 0},
		{`[]`, 0},

		// Part 2 tests.
		{`[1,{"c":"red","b":2},3]`, 4},
		{`{"d":"red","e":[1,2,3,4],"f":5}`, 0},
		{`[1,"red",5]`, 6},
	}

	for _, test := range tests {
		var obj interface{}
		if err := json.Unmarshal([]byte(test.doc), &obj); err != nil {
			t.Fatal(err)
			continue
		}
		if result := AddNumbers(obj); result != test.expected {
			t.Fatalf("For %v, expected %v, got %v", test.doc, test.expected, result)
		}
	}
}
