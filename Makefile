mockery:
	mockery --dir=./service/ --name=AuthService --output=./mocks

test:
	go test -race -coverprofile=coverage.out $(shell go list ./... | grep -v /util/)
	go tool cover -func=coverage.out
	rm coverage.out

lint:
	golangci-lint run ./...

run:
	go run main.go