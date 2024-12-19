package lib

import (
	"bufio"
	"io"
)

func ScanLines(reader io.Reader, receive func(string)) error {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		receive(scanner.Text())
	}
	return scanner.Err()
}
