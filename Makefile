BINARY_NAME=kparser

all: build

build:
	go build -o bin/$(BINARY_NAME) -v .

clean:
	rm -rf bin/$(BINARY_NAME)

run:
	go run main.go
