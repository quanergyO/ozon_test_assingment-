.PHONY: generate, build

generate:
	go run github.com/99designs/gqlgen generate

build:
	go build server.go

run:
	go run server.go

run_memory:
	go run server.go --memory