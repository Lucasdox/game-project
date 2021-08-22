.PHONY:build test

all: build run

build:
	go build -o ./build/bin/game-project ./cmd/game-project

run:
	go run ./cmd/game-project

clean:
	go mod tidy

test:
	go test ./... -cover

migrations:
	migrate -path data/migrations -database "postgresql://root:root@localhost:5432/game?sslmode=disable" -verbose up