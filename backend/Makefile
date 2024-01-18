all: tmp/main

tmp/main: $(shell find . -type f -name '*.go')
	go build -o ./tmp/main ./main.go

clean:
	go clean
	if [ -f tmp/main ] ; then rm tmp/main; fi

run:
	go run .

test:
	go test -race ./...

test\:ci:
	go test -coverprofile=coverage.out -covermode=atomic -race ./...
