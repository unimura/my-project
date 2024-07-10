TARGET := my-project.exe

all: fmt test build

fmt:
	go fmt ./...

test:
	go test ./...

build:
	go build -ldflags="-w -s" -trimpath -buildvcs=false -o $(TARGET)

version:
	go version -m $(TARGET)

run:
	$(TARGET)

clean:
	go clean -i

.PHONY: fmt test build version clean
