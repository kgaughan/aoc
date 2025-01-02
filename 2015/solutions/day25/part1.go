package day25

import "fmt"

const seed = 20151125

func Part1(input string) {
	var row, column int
	_, err := fmt.Sscanf(
		input,
		"To continue, please consult the code grid in the manual.  Enter the code at row %d, column %d.\n",
		&row, &column,
	)
	if err != nil {
		panic(err)
	}

	result := seed
	t := row + column - 2
	for n := (t*(t+1))/2 + column; n > 1; n-- {
		result = (result * 252533) % 33554393
	}
	fmt.Printf("%v\n", result)
}
