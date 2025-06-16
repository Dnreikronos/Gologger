APP_NAME := gologger

.PHONY: build run test clean

build:
	go build -o $(APP_NAME) ./cmd/gologger

run: build
	./$(APP_NAME) -files="testdata/log1.log,testdata/log2.log"

test:
	go test ./...

clean:
	rm -f $(APP_NAME)
