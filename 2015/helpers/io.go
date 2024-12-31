package helpers

import (
	"bufio"
	"io"
)

func ScanLines(reader io.Reader, receive func(string) error) error {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		err := receive(scanner.Text())
		if err != nil {
			return err
		}
	}
	return scanner.Err()
}
