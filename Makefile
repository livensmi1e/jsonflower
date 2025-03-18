.PHONY: build, run

build:
	@go build ./cmd/jsonflower/main.go

run: build
	@./main