package day23

import "fmt"

func Part1(input string) {
	machine := parse(input)
	machine.execute()
	fmt.Printf("%v\n", machine.b)
}
