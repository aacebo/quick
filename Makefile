clean:
	rm -rf ./bin

build:
	go build -o bin/quick src/main.go

clean_build: clean build

run:
	go run ./src/... **/*.q

fmt:
	gofmt -w .

test:
	go clean -testcache
	go test ./... -cover