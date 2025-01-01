package day23

import "fmt"

func Part2(input string) {
	machine := parse(input)
	machine.a = 1
	machine.execute()
	fmt.Printf("%v\n", machine.b)
}
