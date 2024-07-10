.PHONY: build fmt test clean

TARGET := my-project.exe

all: fmt test build

build:
	go build -ldflags="-w -s" -trimpath -buildvcs=false -o $(TARGET)

fmt:
	go fmt ./...

test:
	go test ./...

clean:
	rm $(TARGET)
