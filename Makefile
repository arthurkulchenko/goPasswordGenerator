dev:
	go run main.go

build:
	GOOS=linux GOARCH=amd64 go build -o bin/pass_gen
