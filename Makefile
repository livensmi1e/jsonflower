.PHONY: build, run

build:
	@go build -o ./jsonflower.exe ./cmd/jsonflower/main.go

run: build
	@./jsonflower.exe