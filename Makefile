.PHONY:
.SILENT:
.DEFAULT_GOAL := test

run:
	go run cmd/main.go

tidy:
	GO111MODULE=on go mod tidy

unit-test:
	env GO111MODULE=on go test --short -race -coverprofile=cover.out -v ./...

integration-test:
	newman run postman/api.postman_collection.json

test.coverage:
	env GO111MODULE=on go tool cover -func=cover.out

test: unit-test integration-test test.coverage

lint:
	golangci-lint run

bench:
	env GO111MODULE=on go test -bench=. -cpu=8 -benchmem -cpuprofile=cpu.out -memprofile=mem.out .
