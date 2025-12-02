build:
	go build -o ./bin/main ./src/cmd/commandline/main.go

run:
	go run ./src/cmd/commandline/main.go

test:
	go test ./...
