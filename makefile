run:
	go run ./cmd
wire:
	wire ./internal
build: wire
	go build -o ./bin/app ./cmd 
clean:
	rm -rf ./bin

rr:
	go run -race ./cmd