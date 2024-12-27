package day7

//go:generate goyacc -o parser.go parser.y
//go:generate nex-lexer -o lexer.go lexer.nex

func Part1(input string) {
	Parse(input)
}
