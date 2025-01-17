build:
	go build -o bin/main main.go

run:
	GIN_MODE=release go run .

dev:
	air

test:
	go test -v ./tests/

install:
	go mod download

compile:
	GOOS=linux GOARCH=amd64 go build -o bin/main-linux-amd64 main.go
