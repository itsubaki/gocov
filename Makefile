SHELL := /bin/bash

all:
	make test gocov open
	make install

test:
	go test -cover $(shell go list ./... | grep -v /cmd/ | grep -vxF "$$(go list -m)") -v -coverprofile=coverage.txt -covermode=atomic
	go tool cover -html=coverage.txt -o coverage.html

gocov:
	go run main.go -f coverage.txt -o output/coverage.html

install:
	go install

open:
	open output/coverage.html
