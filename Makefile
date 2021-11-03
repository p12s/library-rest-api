.PHONY:
.SILENT:

test:
	env GO111MODULE=on go test --short -race -coverprofile=cover.out -v ./...
	make test.coverage

test.coverage:
	env GO111MODULE=on go tool cover -func=cover.out

lint:
	golangci-lint run

bench:
	env GO111MODULE=on go test -bench=. -cpu=8 -benchmem -cpuprofile=cpu.out -memprofile=mem.out .
