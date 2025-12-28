.PHONY: run test fmt clean

run:
	go run main.go

build:
	go build -o bin/xray-demo.exe main.go

test:
	go test ./...

fmt:
	go fmt ./...

clean:
	go clean
