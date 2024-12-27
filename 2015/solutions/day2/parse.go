package day2

import (
	"fmt"
	"io"
	"strings"
)

func parse(input string) []dimensions {
	r := strings.NewReader(input)
	result := make([]dimensions, 0, 1000)
	for {
		var d dimensions
		if _, err := fmt.Fscanf(r, "%dx%dx%d\n", &d.l, &d.w, &d.h); err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		result = append(result, d)
	}
	return result
}
