package day8

import (
	"io"
	"strconv"
)

func ParseStrings(reader io.Reader, cb func(string, string)) {
	ScanLines(reader, func(original string) {
		unquoted, _ := strconv.Unquote(original)
		cb(original, unquoted)
	})
}
