package lib

import (
	"strings"
	"testing"
)

type command struct {
	cmd    string
	flag   bool
	c1, c2 Coord
}

type fixture struct {
	script string
	cmds   []command
}

func TestParserHappyPath(t *testing.T) {
	tests := []fixture{
		{"", []command{}},
		{
			`toggle 0,0 through 999,999`,
			[]command{
				{"toggle", false, Coord{0, 0}, Coord{999, 999}},
			},
		},
		{
			`
			turn on 34,76 through 43,87
			turn off 35,77 through 42,86
			toggle 36,78 through 41,85
			`,
			[]command{
				{"turn", true, Coord{34, 76}, Coord{43, 87}},
				{"turn", false, Coord{35, 77}, Coord{42, 86}},
				{"toggle", false, Coord{36, 78}, Coord{41, 85}},
			},
		},
	}

	for _, test := range tests {
		i := 0
		ParseReader(strings.NewReader(test.script), func(cmd string, flag bool, from, to Coord) {
			got := command{cmd, flag, from, to}
			if test.cmds[i] != got {
				t.Errorf(
					"Unexpected result in %q, cmd no. %v: expected %v, got %v",
					test.script, i,
					test.cmds[i], got)
			}
			i++
		})
		if i != len(test.cmds) {
			t.Errorf(
				"Unexpected command count; expected %v, got %v",
				len(test.cmds), i)
		}
	}
}
