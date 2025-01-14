BINARY_NAME=kparser

all: build

build:
	go build -o bin/$(BINARY_NAME) -v .

clean:
	rm -rf bin/$(BINARY_NAME)

run:
# if no args are passed ask for address
ifndef address
	@read -p "Enter address: " address; go run ./... ${address}
else
	go run ./... $(address)
endif
