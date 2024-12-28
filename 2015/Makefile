.PHONY: deps
deps:
	go install golang.org/x/tools/cmd/goyacc@v0.28.0
	go install github.com/blynn/nex@latest

.PHONY: test
test:
	go test ./solutions/... ./helpers/...

.PHONY: generate
generate:
	go generate ./...
