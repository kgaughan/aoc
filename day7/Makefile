build: parser
	go build aoc/day7m

test:
	go test aoc/day7

parser: parser.go lexer.go

parser.go: parser.y
	go tool yacc -o $@ $<

lexer.go: lexer.nex
	nex -o $@ $<

clean:
	rm -f parser.go lexer.go y.output

.PHONY: build clean parser test
