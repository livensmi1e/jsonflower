.PHONY: build, run, test-lex, vtest-lex

build:
	@go build -o ./jsonflower.exe ./cmd/jsonflower/main.go

run: build
	@./jsonflower.exe

test-lex:
	@go test ./internal/lexer

vtest-lex:
	@go test ./internal/lexer/ -v