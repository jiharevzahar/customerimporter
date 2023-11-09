run:
	CSV_PATH=customers/customers.csv go run -race cmd/main/main.go

test:
	go test -race -v ./...

lint:
	golangci-lint run --fix