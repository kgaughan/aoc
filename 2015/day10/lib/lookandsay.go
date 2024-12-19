package lib

import (
	"bytes"
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
		if v, err := readRun(reader); err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		} else {
			buf.WriteString(v.String())
		}
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
		if digit, _, err := s.ReadRune(); err == io.EOF {
			return result, nil
		} else if err != nil {
			return pair{}, err
		} else if digit != result.digit {
			s.UnreadRune()
			return result, nil
		}
		result.count++
	}
}
