.PHONY: run

run:
	go build main.go && ./main

test:
	go test ./...
