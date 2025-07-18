dev:
	go run main.go

build:
	GOOS=linux GOARCH=amd64 go build -o bin/pass_gen

container_build:
	podman build -t pass_gen:latest -f docker/Dockerfile .

container_run:
	podman run -it --name pass_gen pass_gen
