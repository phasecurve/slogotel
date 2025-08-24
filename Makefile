.PHONY: build run clean test

build:
	go build -o bin/slogotel

run:
	go run .

clean:
	rm -rf bin/

test:
	gotestsum --format testname -- -count=1 ./...

tidy:
	go mod tidy