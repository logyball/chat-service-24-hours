test:
	go test -v ./...

bench:
	go test -v -bench=. -benchtime=500x -benchmem ./...

build: test
	go build -o ./bin/server *.go