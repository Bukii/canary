BINARY_NAME=canary

.PHONY: build

build:
	GOARCH=arm GOOS=linux go build -o bin/${BINARY_NAME}-linux_arm main.go
	GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-linux_amd64 main.go

run: build
	./${BINARY_NAME}

clean:
	go clean