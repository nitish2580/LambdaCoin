build:
	go build -o ./bin/WAZIR

run: build
	./bin/WAZIR

test:
	go test -v ./...