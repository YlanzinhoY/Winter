run:
	go run main.go apiExecute

.PHONY: build

build: 
	GOOS=linux GOARCH=amd64 go build -o build/linux/winter
	GOOS=darwin GOARCH=amd64 go build -o build/macos/winter