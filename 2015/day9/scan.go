package day9

import (
	"fmt"
	"io"
)

func ScanLines(reader io.Reader, receive func(string, string, int)) error {
	var from, to string
	var distance int
	for {
		if _, err := fmt.Fscanf(reader, "%s to %s = %d\n", &from, &to, &distance); err == nil {
			receive(from, to, distance)
		} else if err == io.EOF {
			return nil
		} else {
			return err
		}
	}
}
