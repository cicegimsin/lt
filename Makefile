.PHONY: build install clean test

build:
	go build -o lt main.go

install: build
	sudo install -Dm755 lt /usr/local/bin/lt

clean:
	rm -f lt
	go clean

test:
	go test ./...

run:
	go run main.go

fmt:
	go fmt ./...

vet:
	go vet ./...
