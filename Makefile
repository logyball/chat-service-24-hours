test:
	go test -v

build: test
	go build -o ./bin/server *.go