package day12

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
		{`[1,2,3]`, 6},
		{`{"a":2,"b":4}`, 6},
		{`[[[3]]]`, 3},
		{`{"a":{"b":4},"c":-1}`, 3},
		{`{"a":[-1,1]}`, 0},
		{`[-1,{"a":1}]`, 0},
		{`{}`, 0},
		{`[]`, 0},
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
