run: parser
	go run main.go

test:
	go test ./...

parser: lib/parser.go lib/lexer.go

lib/parser.go: lib/parser.y
	goyacc -o $@ $<

lib/lexer.go: lib/lexer.nex
	nex-lexer -o $@ $<

clean:
	rm -f lib/parser.go lib/lexer.go lib/y.output

tools:
	go install golang.org/x/tools/cmd/goyacc@v0.28.0

.PHONY: run clean parser test tools
