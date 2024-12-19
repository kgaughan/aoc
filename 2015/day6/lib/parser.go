package lib

import (
	"fmt"
	"io"
	"strconv"
	"text/scanner"
)

type CmdFunc func(cmd string, flag bool, from, to Coord)

func ParseReader(reader io.Reader, fn CmdFunc) {
	var s scanner.Scanner
	Parse(s.Init(reader), fn)
}

func Parse(s *scanner.Scanner, fn CmdFunc) {
	// This is a pretty dumb parser that assumes that the input is
	// well-formed. The panic() calls are in here just in case it gets
	// something iffy. If this was production code, I'd have better error
	// handling, and likely wouldn't have written my own parser. Also, for
	// my purposes here, Scanner's default error handling works just fine.

	var tok rune
	var flag bool
	for tok != scanner.EOF {
		tok = s.Scan()
		if tok == scanner.EOF {
			break
		}
		cmd := s.TokenText()
		switch cmd {
		case "turn":
			tok = s.Scan()
			text := s.TokenText()
			switch text {
			case "on":
				flag = true
			case "off":
				flag = false
			default:
				panic(fmt.Sprintf("Unexpected token for flag: %q", text))
			}
		case "toggle":
			// Do nothing
		default:
			panic(fmt.Sprintf("Unexpected token for commmand: %q", cmd))
		}

		var c1, c2 Coord

		c1, tok = parseCoordinate(s)
		s.Scan()
		if text := s.TokenText(); text != "through" {
			panic(fmt.Sprintf("Unexpected token; expected 'through': %v", text))
		}
		c2, tok = parseCoordinate(s)

		fn(cmd, flag, c1, c2)
	}
}

func parseCoordinate(s *scanner.Scanner) (Coord, rune) {
	var x, y int

	x, tok := parseInt(s)

	tok = s.Scan()
	text := s.TokenText()
	if text != "," {
		panic(fmt.Sprintf("Expected ',', got: %v", text))
	}

	y, tok = parseInt(s)

	return Coord{x, y}, tok
}

func parseInt(s *scanner.Scanner) (int, rune) {
	tok := s.Scan()
	text := s.TokenText()
	if n, err := strconv.Atoi(text); err != nil {
		panic(err)
	} else {
		return n, tok
	}
}
