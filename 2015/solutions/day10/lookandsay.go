package day10

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
)

type pair struct {
	digit rune
	count int
}

func (p pair) String() string {
	return fmt.Sprintf("%v%c", p.count, p.digit)
}

func LookAndSay(s string) string {
	var buf bytes.Buffer

	reader := strings.NewReader(s)
	for {
		v, err := readRun(reader)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			panic(err)
		}
		buf.WriteString(v.String())
	}

	return buf.String()
}

func readRun(s io.RuneScanner) (pair, error) {
	digit, _, err := s.ReadRune()
	if err != nil {
		return pair{}, err
	}

	result := pair{digit, 1}
	for {
		if digit, _, err := s.ReadRune(); errors.Is(err, io.EOF) {
			return result, nil
		} else if err != nil {
			return pair{}, err
		} else if digit != result.digit {
			err = s.UnreadRune()
			return result, err
		}
		result.count++
	}
}
