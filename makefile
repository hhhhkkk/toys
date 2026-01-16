run:
	go run ./cmd
wire:
	wire ./internal
build: wire
	go build -o ./bin/app ./cmd 