mockery:
	mockery --dir=./service/ --name=AuthService --output=./mocks
	mockery --dir=./service/ --name=UserService --output=./mocks
	mockery --dir=./repository/ --name=UserRepository --output=./mocks
	mockery --dir=./repository/ --name=MenuRepository --output=./mocks
	mockery --dir=./repository/ --name=CouponRepository --output=./mocks

test:
	go test -race -coverprofile=coverage.out $(shell go list ./... | grep -v /util/)
	go tool cover -func=coverage.out
	rm coverage.out

lint:
	golangci-lint run ./...

run:
	go run main.go