runv1:
	go run ./v1/cmd
wirev1:
	wire ./v1/internal
build: wire
	go build -o ./v1/bin/app ./v1/cmd 
clean:
	rm -rf ./v1/bin
rrv1:
	go run -race ./v1/cmd

rr:
	go run -race ./cmd
wirev2:
	wire ./v2/internal